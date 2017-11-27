worker:
	go run cmd/herald-worker/main.go

process:
	go run cmd/herald-worker/main.go
worker: compile
	../../../../bin/herald-worker

compile:
	go install ./...
deps:
	dep ensure
