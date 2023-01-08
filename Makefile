run:
	go run cmd/http/main.go

prometheus:
	docker run \
    -p 9090:9090 \
    --mount source=./dev/prometheus.yml,target=/etc/prometheus/prometheus.yml \
    prom/prometheus