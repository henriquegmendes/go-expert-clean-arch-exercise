deps-up:
	docker-compose up -d

deps-down:
	docker-compose down

start:deps-up
	cd cmd/ordersystem && go run main.go wire_gen.go
