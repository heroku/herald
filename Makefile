run: compile
	../../../../bin/herald

worker: compile
	../../../../bin/herald-worker

compile:
	go install ./...
deps:
	dep ensure
