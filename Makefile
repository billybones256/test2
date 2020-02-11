.PHONY:run
run:
	docker run -p 8080:8081 -it task2

.PHONY:build
build:
	docker build -t task2 .

.PHONY:test
test:
	go test -v

.DEFAULT_GOAL := run