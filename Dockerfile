FROM golang:latest as builder

WORKDIR /go-modules

COPY . ./

# Building using -mod=vendor, which will utilize the v
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod=vendor -o app

#FROM alpine:3.8
FROM scratch

WORKDIR /root/

COPY --from=builder /go-modules/app .

CMD ["./app"]