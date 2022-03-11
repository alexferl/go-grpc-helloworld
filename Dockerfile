FROM golang:1.17.8-alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/server.go

FROM scratch
COPY --from=builder /build/server /app
COPY --from=builder /build/configs /configs

EXPOSE 50051 50052

ENTRYPOINT ["/app"]
CMD ["--env-name", "prod"]
