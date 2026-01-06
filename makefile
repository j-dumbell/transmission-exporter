start-exporter:
	TRANSMISSION_HOST=http://localhost:9091 \
	TRANSMISSION_USER=admin \
	TRANSMISSION_PASSWORD=password \
	PORT=2112 \
	EXPORT_TORRENT_LEVEL_METRICS=true \
	go run ./cmd/exporter/main.go