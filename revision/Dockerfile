FROM golang:1.21.3-alpine3.18 AS builder
RUN apk update && apk add git
#RUN apk add git
ENV GOPROXY=https://goproxy.io,direct
ENV GOPRIVATE=git.mycompany.com,github.com/my/private
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /MultiStage
RUN mkdir /app
WORKDIR /MultiStage
ENV GO111MODULE=on
COPY ./revision ./
#COPY ams_core/go.mod ams_core/go.sum ./
#RUN go mod download
#ADD ./ams_admin_panel_bff ./

#RUN go mod init
RUN go list
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#second stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /MultiStage/main /app/
COPY --from=builder /MultiStage/.env /app/

EXPOSE 5001
## Our start command which kicks off
## our newly created binary executable
CMD "/app/main"