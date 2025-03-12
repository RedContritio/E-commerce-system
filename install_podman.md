从源码安装 podman 似乎很麻烦，在网络较差的情况下，所以我改用 brew

```bash
# install brew
export HOMEBREW_BREW_GIT_REMOTE="https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git"
export HOMEBREW_API_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles/api"
## 变量名：https://github.com/Homebrew/brew/blob/ea5f6dacd3a5f263ec4bccde2a6297ddff6b1c1a/docs/Tips-N'-Tricks.md?plain=1#L134
## 换源目标：https://blog.51cto.com/u_13162375/13264759
export HOMEBREW_ARTIFACT_DOMAIN="https://ghcr.nju.edu.cn"

git clone --depth=1 https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/install.git brew-install
/bin/bash brew-install/install.sh
rm -rf brew-install/

# add brew into PATH
test -d ~/.linuxbrew && eval "$(~/.linuxbrew/bin/brew shellenv)"
test -d /home/linuxbrew/.linuxbrew && eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
test -r ~/.bash_profile && echo "eval \"\$($(brew --prefix)/bin/brew shellenv)\"" >> ~/.bash_profile
test -r ~/.profile && echo "eval \"\$($(brew --prefix)/bin/brew shellenv)\"" >> ~/.profile
test -r ~/.zprofile && echo "eval \"\$($(brew --prefix)/bin/brew shellenv)\"" >> ~/.zprofile

# install podman-5.4.1 (latest)
brew install podman
brew services start podman

brew install podman-compose

# podman need cgroup2
# https://github.com/spurin/wsl-cgroupsv2
```