#!/bin/bash

set -e

os=$(uname)
arch=$(uname -m)

case $arch in
    aarch64|arm64)
        arch="arm64" 
        ;;
    x86_64|amd64)
        arch="amd64"
        ;;
    *)
        echo "$arch not supported by this script"
        exit 1
        ;;
esac

curl -L -o gitviz "https://github.com/riezebosch/gitviz/releases/latest/download/gitviz_${os}_${arch}"
chmod +x gitviz
sudo mv gitviz /usr/local/bin/gitviz
