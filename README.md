## gitviz

![test](https://github.com/riezebosch/gitviz/workflows/test/badge.svg)

`gitviz` is a tool designed for educational purposes to visualize the Git repository in real time, providing a graphical representation of the blobs, trees, and commits that make up the [git data structure](https://eagain.net/articles/git-for-computer-scientists/).

[![graph](img/graph.png)https://riezebosch.github.io/gitviz/](https://riezebosch.github.io/gitviz/)

### Key Features

- Real-time visualization of Git repository
- Graphical representation of blobs, trees, commits, and references (branches, tags, or HEAD)
- An educational tool for understanding Git's data structure

### Installation

You can install `gitviz` by following these steps:

1. Open a terminal and navigate to the root of your repository.

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

Alternatively, you can run the following command to install `gitviz` directly:

```sh
curl -sSL https://riezebosch.github.io/gitviz/install.sh | sh
```

### Usage

To use `gitviz`, navigate to the root of your repository in a terminal and run the following command:

```sh
gitviz
```

and use `CTRL+C` to stop it when done.

You can also run `gitviz` in the background by appending an ampersand (`&`) to the command:

```sh
gitviz &
```

To stop the process when running in the background, you can use the `kill` command with the appropriate job number. For example, if the job number is `1`, you can use the following command:

```sh
kill %1
```

By using `gitviz`, you will gain valuable insights into the structure of your Git repository in real time, with different colors representing commits, trees, blobs, and references.

---

This [README.md](./README.md) was peer-reviewed by [ChatGPT](https://chat.openai.com/).