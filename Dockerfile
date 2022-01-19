FROM golang:1.17-alpine3.15 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o bogops_exporter .

FROM alpine:3.15
COPY --from=builder /build/bogops_exporter /
EXPOSE 8080
ENTRYPOINT ["/bogops_exporter"]
