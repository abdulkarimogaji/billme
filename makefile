server:
	go run main.go

migrate-up:
	migrate -path db/migration -database "mysql://root@tcp(127.0.0.1:3306)/billme?checkConnLiveness=false&maxAllowedPacket=0&parseTime=true" -verbose up

migrate-down:
	migrate -path db/migration -database "mysql://root@tcp(127.0.0.1:3306)/billme?checkConnLiveness=false&maxAllowedPacket=0&parseTime=true" -verbose down

.PHONY: server migrate-up migrate-down