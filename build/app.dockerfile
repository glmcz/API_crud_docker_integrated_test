FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && mkdir /app

WORKDIR /app
ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
COPY ./build/app-linux app
COPY ./config/configDocker.yaml ./config/config.yaml

COPY ./static/index.html ./static/index.html
COPY ./static/thread.html ./static/thread.html

COPY ./build/startup.sh /startup.sh
ENTRYPOINT ["sh","/startup.sh"]

