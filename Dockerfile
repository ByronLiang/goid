FROM alpine:3.10

RUN apk update && apk add curl && rm -rf /var/cache/apk/*

WORKDIR /app

COPY main ./

EXPOSE 8090
EXPOSE 9090

CMD ["./main"]
