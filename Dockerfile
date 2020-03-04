ARG VERSION=latest

FROM golang:$VERSION as builder

RUN mkdir -p /gomod/golang-api
WORKDIR /gomod/golang-api
ADD /go.mod go.mod
ADD /go.sum go.sum
ADD /Makefile Makefile
ADD /.env .env
ADD /person-post person-post
WORKDIR /gomod/golang-api/person-post
RUN make test 
RUN make build
RUN ls -la /gomod/golang-api/person-post/bin

## ----------------------------------------------------------------------------
FROM scratch
WORKDIR /

COPY --from=builder /gomod/golang-api/person-post/bin/v1 /golang-api/person-post/v1
COPY --from=builder /gomod/golang-api/.env /.env

ENTRYPOINT [ "/golang-api/person-post/v1" ]
