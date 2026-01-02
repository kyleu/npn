#!/bin/bash

## Copies build files to main project

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build.sh

mkdir -p ../web/assets/vendor

rm -f ../assets/vendor/*.js
rm -f ../assets/vendor/*.css
rm -f ../assets/vendor/*.map

cp dist/assets/index-*.js ../assets/vendor/npn.js
cp dist/assets/index-*.css ../assets/vendor/npn.css
cp dist/vendor/editor/editor.js ../assets/vendor/vendor.js
