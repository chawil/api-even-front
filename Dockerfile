FROM golang:alpine AS dependencies
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download -x

FROM dependencies AS builder
COPY . .
RUN go build -v -o web

FROM alpine AS final
COPY --from=builder /app/web .
ENV GIN_MODE=release
ENTRYPOINT ["/web"]
