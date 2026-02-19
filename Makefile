clean:
	go clean

web:
	CGO_ENABLED=0 go build -v -o bin ./cmd/web/main.go

dev:
	docker compose -f ./deploy/dev/compose.yml up --build
