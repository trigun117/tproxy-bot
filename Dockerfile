FROM alpine

COPY tproxy-bot .

RUN apk add --no-cache ca-certificates &&\
    chmod +x tproxy-bot

ENTRYPOINT [ "./tproxy-bot" ]