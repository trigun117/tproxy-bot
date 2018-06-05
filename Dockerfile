FROM alpine

COPY tproxy-bot .

RUN apk update && apk add --no-cache ca-certificates

ENTRYPOINT [ "./tproxy-bot" ]