up:
	docker-compose up

down:
	docker-compose down -v
	
build-image:
	docker build -f go/docker/Dockerfile.dockService -t smakamura/test .