FROM golang:1.14.1
ENV GOPROXY "https://goproxy.cn"
WORKDIR $GOPATH/src/yoyo-mall
COPY . .
RUN go build -o main -v .
EXPOSE 4096
CMD ["./main", "-c", "config/pre.yaml"]
