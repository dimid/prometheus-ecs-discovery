

build:
	docker build -t prometheus-ecs-discovery:latest .

test:
	go test -v .
