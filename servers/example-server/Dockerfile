# servers/example-server/Dockerfile
FROM golang:1.20-alpine AS build
WORKDIR /src
COPY . .
RUN apk add --no-cache git && go build -o /out/example-server ./cmd/server

FROM alpine:3.18
RUN addgroup -S app && adduser -S -G app app
COPY --from=build /out/example-server /usr/local/bin/example-server
USER app
EXPOSE 8080
HEALTHCHECK --interval=10s --timeout=3s --retries=3 CMD wget -qO- http://127.0.0.1:8080/health/ready || exit 1
ENTRYPOINT ["/usr/local/bin/example-server"]
