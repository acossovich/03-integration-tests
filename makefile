test:
	go	test	./...	-v
coverage:
	go	test	./...	-v	-coverprofile=coverage.out
	go tool cover -html=coverage.out
linter:
	golangci-lint run