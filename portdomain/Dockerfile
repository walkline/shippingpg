FROM golang:alpine as builder
COPY ./ /go/src/github.com/walkline/shippingpg/portdomain
WORKDIR /go/src/github.com/walkline/shippingpg/portdomain/cmd/portdomainservice
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/walkline/shippingpg/portdomain/cmd/portdomainservice /app/
ENTRYPOINT ["./portdomainservice"]
