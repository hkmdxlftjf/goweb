FROM golang:1.18 as builder

WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download
RUN CGO_ENABLED=0 go build -o gin-web

FROM alpine:3.10
COPY --from=builder /app/gin-web /
CMD ["/gin-web"]
