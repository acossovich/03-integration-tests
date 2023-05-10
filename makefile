testeo_gral:
	go	test	./...	-v	-cover
linter:
	golangci-lint run