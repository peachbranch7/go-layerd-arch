FROM golang:latest AS build

ENV GO111MODULE=on

WORKDIR /go/src/github.com/peachbranch7/go-layered-arch

COPY . .

# RUN rm -rf /go/pkg/mod && go mod tidy
RUN rm -rf /go/pkg/mod && go mod download
RUN apt-get update && apt-get install -y git
# RUN apk update && apk add --no-cache git
RUN cd cmd && go build -v -o main

FROM alpine:3.8

COPY --from=build /go/src/github.com/peachbranch7/go-layered-arch/cmd/main /usr/local/bin/sample

CMD ["sample"]