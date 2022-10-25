FROM alpine:edge

RUN apk add --update ca-certificates
WORKDIR /root

COPY ../build/linux-amd64/bzed /usr/bin/bzed

CMD ["bzed", "start"]