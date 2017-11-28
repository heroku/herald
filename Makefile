worker: 
	go run cmd/herald-worker/main.go

proc: 
	go run cmd/herald/main.go 

compile:
	go install ./...
deps:
	dep ensure
