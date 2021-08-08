
#build stage
FROM golang:alpine3.14 AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN ls
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:3.14
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=Txy_server_0.01 Version=0.0.1
EXPOSE 5000
RUN pwd
CMD [ "Txy_server" ]