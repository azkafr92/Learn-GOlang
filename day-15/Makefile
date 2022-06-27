IMAGE_NAME=alterra-tech-academy-day15-app
CONTAINER_NAME=alterra-tech-academy-day15-app1
TAG=latest
REPOSITORY=azkafr92
run:
	cd app && \
	go run main.go

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-create:
	docker create -p 8080:8080 --name ${CONTAINER_NAME} ${IMAGE_NAME}

docker-start:
	docker start ${CONTAINER_NAME}

docker-stop:
	docker stop ${CONTAINER_NAME}

docker-push:
	docker push ${REPOSITORY}/${IMAGE_NAME}:${TAG}