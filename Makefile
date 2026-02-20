clean:
	go clean

web:
	bun build ./script/index.js --bundle --outfile=./static/index.js
	bun x sass ./style/main.scss ./static/style.css
	CGO_ENABLED=0 go build -v -o bin ./cmd/web/main.go

dev:
	docker compose -f ./deploy/dev/compose.yml up --build
