FROM golang:alpine as builder
COPY ./ /go/src/github.com/walkline/shippingpg/clientapi
WORKDIR /go/src/github.com/walkline/shippingpg/clientapi/cmd/clientapi
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/walkline/shippingpg/clientapi/cmd/clientapi /app/
ENTRYPOINT ["./clientapi"]
