FROM golang:1.13-alpine

WORKDIR /work_api

RUN apk add --no-cache alpine-sdk git && go get -u github.com/oxequa/realize

EXPOSE 8082
CMD ["realize", "start"]