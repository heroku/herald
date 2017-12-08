worker: 
	go run cmd/version-scraper/main.go

bin-worker: compile
	../../../../bin/version-scraper

compile:
	go install ./...

deps:
	dep ensure
