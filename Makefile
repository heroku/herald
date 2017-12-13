worker:
	go run cmd/version-scraper/main.go

bin-worker: compile
	../../../../bin/version-scraper

test:
	go run cmd/test/main.go

web:
	go run cmd/web/main.go

compile:
	go install ./...

deps:
	dep ensure
