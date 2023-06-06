case $(uname -m) in
    aarch64|arm64)
    curl -L -o gitviz "https://github.com/riezebosch/gitviz/releases/latest/download/gitviz_$(uname)_arm64" 
    ;; amd64)
    curl -L -o gitviz "https://github.com/riezebosch/gitviz/releases/latest/download/gitviz_$(uname)_amd64"
    ;; *) echo "$(uname -m) not supported by this script" exit 1
esac

chmod +x gitviz
sudo mv gitviz /usr/local/bin/gitviz