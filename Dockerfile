FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/bytegriffin/go_api_server
COPY . $GOPATH/src/github.com/bytegriffin/go_api_server
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go_api_server"]