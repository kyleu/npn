#!/bin/bash

## Builds desktop release artifacts for macOS, Linux, and Windows.
##
## Usage:
##   ./bin/build/desktop.release.sh [version] [-y|--yes]
##
## Arguments:
##   version  Version tag for output filenames (default: v0.0.0).
##
## Requires:
##   - Docker
##   - appdmg, codesign, lipo (macOS)
##   - APPLE_SIGNING_IDENTITY set for codesign
##   - curl and zip
##
## Outputs:
##   - build/dist/npn_<version>_darwin_*_desktop.(dmg|zip)
##   - build/dist/npn_<version>_linux_amd64_desktop.zip
##   - build/dist/npn_<version>_windows_amd64_desktop.zip

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "error: required command '$1' not found${2:+ ($2)}" >&2
    exit 1
  fi
}

require_env() {
  if [[ -z "${!1:-}" ]]; then
    echo "error: required environment variable '$1' is not set" >&2
    exit 1
  fi
}

require_cmd docker "install Docker Desktop from https://www.docker.com/products/docker-desktop/"
require_cmd zip "install zip from your package manager"
require_cmd curl "install curl from your package manager"
require_cmd appdmg "install via npm i -g appdmg"
require_cmd codesign "requires macOS codesign tool"
require_cmd lipo "requires macOS Xcode command line tools"
require_env APPLE_SIGNING_IDENTITY

TGT=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --) shift; break;;
    -*)
      echo "unknown option: $1" >&2
      exit 1
      ;;
    *)
      if [[ -z "$TGT" ]]; then
        TGT="$1"
        shift
      else
        echo "unexpected argument: $1" >&2
        exit 1
      fi
      ;;
  esac
done

TGT=${TGT:-v0.0.0}

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
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./NPN.app/Contents/MacOS/npn"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./NPN.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./npn_${TGT}_darwin_amd64_desktop.dmg"
zip -r "npn_${TGT}_darwin_amd64_desktop.zip" "./NPN.app"

# darwin arm64
cp "npn.darwin.arm64" "./NPN.app/Contents/MacOS/npn"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./NPN.app/Contents/MacOS/npn"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./NPN.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./npn_${TGT}_darwin_arm64_desktop.dmg"
zip -r "npn_${TGT}_darwin_arm64_desktop.zip" "./NPN.app"

# macOS universal
rm "./NPN.app/Contents/MacOS/npn"
lipo -create -output "./NPN.app/Contents/MacOS/npn" npn.darwin npn.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./NPN.app/Contents/MacOS/npn"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./NPN.app"

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

cd "$dir/../.."
echo "Builds written to ./build/dist (npn_${TGT}_*_desktop.*)"
