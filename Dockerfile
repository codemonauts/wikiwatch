# argument for Go version

FROM golang:1.17-alpine AS builder
RUN apk add --no-cache git
WORKDIR /src
COPY ./ ./
RUN CGO_ENABLED=0 go build

FROM gcr.io/distroless/static

LABEL maintainer="Felix Breidenstein (felix@codemonauts.com)"
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /src/wikiwatch /wikiwatch

ENTRYPOINT ["/wikiwatch","-config", "/config.json"]
