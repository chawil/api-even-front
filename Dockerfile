FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o web

FROM alpine
COPY --from=0 /app/web .
ENTRYPOINT ["/web"]
