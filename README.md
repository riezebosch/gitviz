## gitviz

![test](https://github.com/riezebosch/gitviz/workflows/test/badge.svg)

For education purposes: visualize the git repository in real-time revealing the blobs, trees and commits that make up the [git data structure](https://eagain.net/articles/git-for-computer-scientists/).

![graph](img/graph.png)

color | type
------|-------
red   | commit
green | tree
blue  | blob
grey  | ref (branch, tag or HEAD)

## install

Download from [releases](../../releases) and start from the root of your repo.

```sh
case $(uname -m) in
    aarch64|arm64)
    curl -L -o gitviz "https://github.com/riezebosch/gitviz/releases/latest/download/gitviz_$(uname)_arm64" 
    ;; amd64)
    curl -L -o gitviz "https://github.com/riezebosch/gitviz/releases/latest/download/gitviz_$(uname)_amd64"
    ;; *) echo "$(uname -m) not supported by this script" exit 1
esac

chmod +x gitviz
sudo mv gitviz /usr/local/bin/gitviz
```

or run: 

```sh
wget -O - https://riezebosch.github.io/gitviz/install.sh | sh
```

run from the root of your repository:

```sh
gitviz
```

or 

```sh
gitviz &
...
kill %1
```

to push the process to the background and stop it when finished.