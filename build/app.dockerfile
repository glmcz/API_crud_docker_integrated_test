FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && mkdir /app

WORKDIR /app

COPY ./build/app-linux app
COPY ./config/configDocker.yaml ./config/config.yaml

CMD ["/app/app"]
