FROM alpine:3.10

WORKDIR /app

COPY main ./

EXPOSE 8090
EXPOSE 9090

CMD ["./main"]
