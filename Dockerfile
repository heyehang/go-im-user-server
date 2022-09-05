FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct


WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY ./etc /app/etc
COPY ./user_server.pb /app/user_server.pb
RUN go build -ldflags="-s -w" -o /app/main main.go


FROM golang

COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=golang /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/user_server.pb /app/user_server.pb

CMD ["./main", "-f", "etc/imserver.yaml",, "-gateway", "etc/gateway.yaml"]