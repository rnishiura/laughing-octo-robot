# FROM golang:1.19.1-alpine
 
# RUN apk update && \
#     apk upgrade && \
#     apk add git bash

# WORKDIR /go/src/api
# # COPY api/go.mod api/go.sum ./
# RUN go mod init auth-api
# # ENV GOPATH=
# RUN go get github.com/cespare/reflex
# RUN go get -u github.com/gofiber/fiber/v2
# ENV CGO_ENABLED=0
 
# WORKDIR /go/src/api
# COPY ./api ./

# RUN go mod download


FROM golang:1.18.1-alpine
 
RUN apk update && \
    apk upgrade && \
    apk add git bash

WORKDIR /go/src/api
COPY api/go.mod api/go.sum ./
# RUN go mod init auth-api
RUN go get -u github.com/gofiber/fiber/v2
RUN go install github.com/cosmtrek/air@latest
# RUN go get github.com/cespare/reflex
# CMD ["air"]
ENV CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
 
WORKDIR /go/src/api
COPY ./api ./
 
RUN go mod download