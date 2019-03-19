ARG GO_VERSION=1.11.5

FROM golang:${GO_VERSION} AS builder
ENV GOPATH /usr
ENV APP ${GOPATH}/src/github.com/joshiomkarj/inMemoryServer/app
COPY /app ${APP}/
WORKDIR ${APP}/cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app

FROM golang:1.11-alpine
WORKDIR /
COPY --from=builder /app /
ENTRYPOINT ["/app"]
EXPOSE 8080