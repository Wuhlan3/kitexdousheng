FROM golang:latest
WORKDIR /kitexdousheng
ADD . /kitexdousheng
ENV GOPROXY https://goproxy.cn
EXPOSE 8081
CMD go mod tidy
CMD cd /kitexdousheng/cmd/api && go run main.go router.go