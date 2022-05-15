# syntax=docker/dockerfile:1.4

FROM golang:1.18 AS build-dev
WORKDIR /go/src/app
COPY --link go.mod go.sum ./
RUN go mod download
COPY --link . .
RUN go install -buildvcs=false -trimpath -ldflags '-w -s' -v
FROM debian:buster-slim AS stage
COPY --from=build-dev /go/bin/entgo-bbs /go/bin/entgo-bbs
VOLUME ["/data"]
CMD ["/go/bin/entgo-bbs"]
