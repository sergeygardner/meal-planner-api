FROM golang:alpine
LABEL remove=true mantainer=sergeygardner@gmail.com

ARG PROJECT_ENV=prod
ENV PROJECT_ENV=$PROJECT_ENV

COPY ./app/ /go/src/app
COPY ./service/run/ /run/

RUN apk upgrade && apk update && apk add git protoc && cd /go/src/app && go get -d -v ./... && go install -v ./... \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

WORKDIR /go/src/app

CMD ["sh"]