#FROM golang:1.16-alpine
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod dowload
#
#COPY *.go ./
#
#RUN go build -o /dockerfile
#
#EXPOSE 8080
#
#CMD ["/dockerfile"]

FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod dowload

COPY *.go ./
RUN go build -v /
#deploy

FROM alpine:3.9
WORKDIR /root/
COPY --from=builder /app .
ENTRYPOINT ["./dockerfile"]
EXPOSE 80

CMD ["/dockerfile"]