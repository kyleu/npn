#!/bin/bash

## Meant to be run as part of the release process, builds desktop apps

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

if command -v retry &> /dev/null
then
  retry -t 4 -- docker build -f tools/desktop/Dockerfile.desktop -t npn .
else
  docker build -f tools/desktop/Dockerfile.desktop -t npn .
fi


rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create npn)
docker cp "$id":/dist - > ./desktop.tar
docker rm -v "$id"
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/npn "npn.darwin"
mv dist/darwin_arm64/npn "npn.darwin.arm64"
mv dist/linux_amd64/npn "npn"
mv dist/windows_amd64/npn "npn.exe"
rm -rf "dist"

# darwin amd64
cp -R "../../tools/desktop/template" .

mkdir -p "./NPN.app/Contents/Resources"
mkdir -p "./NPN.app/Contents/MacOS"

cp -R "./template/darwin/Info.plist" "./NPN.app/Contents/Info.plist"
cp -R "./template/darwin/icons.icns" "./NPN.app/Contents/Resources/icons.icns"

cp "npn.darwin" "./NPN.app/Contents/MacOS/npn"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./NPN.app/Contents/MacOS/npn"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./NPN.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./npn_${TGT}_darwin_amd64_desktop.dmg"
zip -r "npn_${TGT}_darwin_amd64_desktop.zip" "./NPN.app"

# darwin arm64
cp "npn.darwin.arm64" "./NPN.app/Contents/MacOS/npn"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./NPN.app/Contents/MacOS/npn"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./NPN.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./npn_${TGT}_darwin_arm64_desktop.dmg"
zip -r "npn_${TGT}_darwin_arm64_desktop.zip" "./NPN.app"

# macOS universal
rm "./NPN.app/Contents/MacOS/npn"
lipo -create -output "./NPN.app/Contents/MacOS/npn" npn.darwin npn.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./NPN.app/Contents/MacOS/npn"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./NPN.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./npn_${TGT}_darwin_all_desktop.dmg"
zip -r "npn_${TGT}_darwin_all_desktop.zip" "./NPN.app"

# linux
echo "building Linux zip..."
zip "npn_${TGT}_linux_amd64_desktop.zip" "./npn"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "npn_${TGT}_windows_amd64_desktop.zip" "./npn.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./npn_${TGT}_darwin_amd64_desktop.dmg" "../../build/dist"
mv "./npn_${TGT}_darwin_amd64_desktop.zip" "../../build/dist"
mv "./npn_${TGT}_darwin_arm64_desktop.dmg" "../../build/dist"
mv "./npn_${TGT}_darwin_arm64_desktop.zip" "../../build/dist"
mv "./npn_${TGT}_darwin_all_desktop.dmg" "../../build/dist"
mv "./npn_${TGT}_darwin_all_desktop.zip" "../../build/dist"
mv "./npn_${TGT}_linux_amd64_desktop.zip" "../../build/dist"
mv "./npn_${TGT}_windows_amd64_desktop.zip" "../../build/dist"
