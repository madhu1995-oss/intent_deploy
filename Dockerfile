FROM golang:1.8-alpine
ADD . /go/src/intent_deploy
#RUN go install intent_deploy

FROM alpine:latest
COPY --from=0 /go/bin/intent_deploy  .
ENV PORT 8080
CMD ["./intent_deploy"]
