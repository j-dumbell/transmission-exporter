start-exporter:
	TRANSMISSION_HOST=http://localhost:9091 \
	TRANSMISSION_USER=admin \
	TRANSMISSION_PASSWORD=password \
	PORT=2112 \
	EXPORT_TORRENT_LEVEL_METRICS=true \
	go run ./cmd/exporter/main.go

build-exporter:
	go build ./cmd/exporter

docker-build-exporter:
	docker build -f cmd/exporter/Dockerfile .

fmt:
	gofmt -w .

vet:
	go vet ./...

test:
	go test ./... -race
