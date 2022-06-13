FROM golang:1.18 as builder

RUN mkdir -p /go/src/github.com/aveplen-bach/config-gateway

WORKDIR /go/src/github.com/aveplen-bach/config-gateway

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /bin/config-gateway \
    /go/src/github.com/aveplen-bach/config-gateway/cmd/main.go

FROM alpine:3.15.4 as runtime

COPY --from=builder /bin/config-gateway /bin/config-gateway
COPY ./config-gateway.yaml ./config-gateway.yaml

ENTRYPOINT [ "/bin/config-gateway" ]