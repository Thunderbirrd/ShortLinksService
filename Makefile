gen-mocks:
	mockgen -source=internal/repository/repository.go -destination=internal/mocks/mock_repository/mock_repository.go

build-image:
	docker build -t links-service .

docker-start:
	docker run --name=links-service --publish 8080:8080 --rm links-service