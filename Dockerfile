FROM golang:alpine

WORKDIR /build

COPY . .

EXPOSE 9090

RUN go env -w GOPROXY=https://goproxy.cn && go build -o myjwt main.go 

CMD ["./myjwt"]
