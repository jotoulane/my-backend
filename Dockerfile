FROM golang:1.21-alpine

WORKDIR /src
COPY . .

ENV GOPROXY https://goproxy.cn
RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o blog-backend

ENTRYPOINT [ "./blog-backend" ]
