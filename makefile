start-exporter:
	TRANSMISSION_HOST=http://localhost:9091 \
	TRANSMISSION_USER=admin \
	TRANSMISSION_PASSWORD=password \
	PORT=2112 \
	EXPORT_TORRENT_LEVEL_METRICS=true \
	go run ./cmd/exporter/main.go

build-exporter:
	go build ./cmd/exporter

DOCKER_USERNAME ?= jdumbell92
IMAGE_NAME ?= transmission-exporter
IMAGE_TAG ?= latest
IMAGE ?= $(DOCKER_USERNAME)/$(IMAGE_NAME):$(IMAGE_TAG)

docker-build-exporter:
	docker build -f cmd/exporter/Dockerfile -t $(IMAGE) .

docker-push-exporter: docker-build-exporter
	docker push $(IMAGE)

docker-build-push-exporter: docker-push-exporter

fmt:
	gofmt -w .

vet:
	go vet ./...

test:
	go test ./... -race
