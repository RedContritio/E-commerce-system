# E-commerce-system
E-commerce system as a practice

# Dependencies

```bash
## for WSL
sudo mount --make-shared /

docker pull golang:1.24-alpine postgres:17-alpine

docker run -d --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=ecommerce -p 5432:5432 postgres:17-alpine

# if port is used
# sudo lsof -i :<port> | awk 'NR!=1 {print $2}' | xargs sudo kill -9
# docker rm -f postgres

docker run -it --rm -v $(pwd):/app -w /app -p 8080:8080 -e GOPROXY=https://goproxy.cn golang:1.24-alpine sh
```