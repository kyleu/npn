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

cp dist/js/app.*.js ../assets/vendor/npn.js
cp dist/js/chunk*.*.js ../assets/vendor/vendor.js
cp dist/js/*.map ../assets/vendor
cp dist/css/app.*.css ../assets/vendor/npn.css
