gen-mocks:
	mockgen -source=internal/repository/repository.go -destination=internal/mocks/mock_repository/mock_repository.go