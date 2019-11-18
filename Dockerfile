FROM golang:1.8-alpine
ADD . /go/src/intent_deploy

FROM alpine:latest
ENV PORT 8080
CMD ["./intent_deploy"]
