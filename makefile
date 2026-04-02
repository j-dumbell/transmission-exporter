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
PLATFORMS ?= linux/amd64,linux/arm64

docker-buildx-setup:
	docker buildx inspect multiplatform > /dev/null 2>&1 || docker buildx create --name multiplatform --use --bootstrap

docker-build-exporter: docker-buildx-setup
	docker buildx build --platform $(PLATFORMS) -f cmd/exporter/Dockerfile -t $(IMAGE) .

docker-push-exporter: docker-buildx-setup
	docker buildx build --platform $(PLATFORMS) -f cmd/exporter/Dockerfile -t $(IMAGE) --push .

docker-build-push-exporter: docker-push-exporter

fmt:
	gofmt -w .

vet:
	go vet ./...

test:
	go test ./... -race
