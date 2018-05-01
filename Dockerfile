FROM alpine

COPY tproxy-bot .

RUN apk update &&\
    apk add --no-cache ca-certificates
#    apk add --no-cache ca-certificates openssl

ENTRYPOINT [ "./tproxy-bot" ]