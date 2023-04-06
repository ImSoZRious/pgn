FROM golang:1.20-alpine as builder

WORKDIR /build

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main ./...

FROM alpine as master

WORKDIR /app

ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
ENV GRPC_GO_LOG_SEVERITY_LEVEL=info

COPY --from=builder /build/main .

CMD ["./main"]
