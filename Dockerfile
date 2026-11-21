# Build: docker build --build-arg VERSION=v1.4.0 -t simple-proxy .
# Run:   docker run -p 8888:8888 simple-proxy [options]
FROM golang:1.27-alpine AS builder

ARG VERSION=development

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X 'main.Version=${VERSION}'" -o /out/simple-proxy .


FROM alpine:3.24

# Pick up Alpine security fixes published since the base image was built,
# and run as an unprivileged user.
RUN apk upgrade --no-cache \
    && adduser -D -H -u 10001 proxy

COPY --from=builder /out/simple-proxy /usr/local/bin/simple-proxy
COPY LICENSE /usr/share/licenses/simple-proxy/LICENSE

USER proxy
EXPOSE 8888

ENTRYPOINT ["simple-proxy"]
