FROM golang:latest
WORKDIR /go/src/app

COPY bin/cmd .

CMD ["cmd"]