FROM golang:alpine AS builder

# 移动到工作目录 /app
WORKDIR /app

#为镜像设置环境变量
ENV Go111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go generate -x ./... && go build -o bin/main main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/migrate .
COPY --from=builder /app/cert/ ./cert
COPY --from=builder /app/static/ ./static
COPY --from=builder /app/bin/main .
COPY --from=builder /app/internal/dao/db/migration ./migration
COPY --from=builder /app/start.sh .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/config ./config

RUN chmod +x wait-for.sh
RUN chmod +x start.sh

EXPOSE 8080

#ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main","-path=/app/config/app_docker"]
