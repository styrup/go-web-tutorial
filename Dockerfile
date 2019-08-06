FROM golang:alpine AS builder
RUN apk --update add git
ADD . /src
WORKDIR /src
RUN go get -v -t -d ./...
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /src/main /root/
CMD ["/root/main"]