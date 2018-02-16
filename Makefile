test:
	go run cmd/herald-test/main.go
run:
	go run cmd/version-scraper/main.go

bin-worker: compile
	../../../../bin/version-scraper

cli:
	go run cmd/herald-cli/main.go

web:
	go run cmd/web/main.go

compile:
	go install ./...

deps:
	dep ensure
