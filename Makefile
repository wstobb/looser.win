clean:
	go clean
	rm -rf ./cmd/web/static
	rm -rf ./cmd/web/templates

web: clean
	bun build ./script/index.js --bundle --outfile=./cmd/web/static/index.js
	bun x sass ./style/main.scss ./cmd/web/static/style.css
	cp -r ./templates ./cmd/web/
	CGO_ENABLED=0 go build -v -o bin ./cmd/web/main.go

dev:
	docker compose -f ./deploy/dev/compose.yml up --build
