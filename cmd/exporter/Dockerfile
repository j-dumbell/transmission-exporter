FROM --platform=$BUILDPLATFORM golang:1.25.5-bookworm AS builder
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/exporter

FROM gcr.io/distroless/static-debian12
COPY --from=builder /out/app /app
USER 65532:65532
ENTRYPOINT ["/app"]
