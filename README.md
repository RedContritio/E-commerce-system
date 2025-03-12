# E-commerce-system
E-commerce system as a practice

# Dependencies

```bash
apt-get install podman-docker
## for WSL
sudo mount --make-shared /

## Ensure podman have registry
## /etc/containers/registries.conf
##
## unqualified-search-registries = ["docker.io"]
##

podman pull golang:1.24-alpine

podman run -it --rm -v $(pwd):/app -w /app -p 8080:8080 golang:1.24-alpine sh
```