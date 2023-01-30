##################################
# Golang Environment Image: 1.19 #
# -> Build runnable file         #
##################################
FROM golang:1.19 as stage

RUN mkdir /build && \
    chmod -R 777 /build

COPY . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o job ./cmd/main.go

########################################################################
# Standard Alpine image: latest                                        #
# -> Copy ca-certs                                                     #
# -> Credits: https://github.com/drone/ca-certs/blob/master/Dockerfile #
########################################################################
FROM alpine:latest as alpine

RUN apk add -U --no-cache ca-certificates

##################################
# Scratch Image                  #
# -> Empty image to execute app  #
##################################
FROM scratch as runner

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=stage /build/job job
COPY --from=stage /build/cmd/env/.env.json .env.json

CMD ["./job"]


