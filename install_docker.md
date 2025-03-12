
```bash
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings

# 官方源和密钥，很慢，不建议使用
## 添加密钥
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
## 添加源
echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu   "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 阿里云源，建议使用
# curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/ali-docker.gpg
sudo chmod a+r /etc/apt/keyrings/ali-docker.gpg
echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/ali-docker.gpg] https://mirrors.aliyun.com/docker-ce/linux/ubuntu   "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/ali-docker.list > /dev/null

sudo apt-get update

# install docker
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-compose
```

change /etc/docker/daemon.json to 

```json
{
    "auths": { }, 
    "registry-mirrors": [
        "https://docker.mirrors.ustc.edu.cn", 
        "https://docker.tbedu.top"
    ]
}
```

```bash
# apply mirrors
sudo systemctl daemon-reload && sudo systemctl restart docker.service

# hello-world
sudo docker run hello-world
```