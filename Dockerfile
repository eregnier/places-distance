FROM alpine:latest as alpine
RUN apk add -U --no-cache ca-certificates

WORKDIR /

COPY app .
EXPOSE  8080

ENTRYPOINT [ "/app" ]