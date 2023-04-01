.PHONY: gen-users-json

gen-users-json:
	@echo "Generating users.json"
	@go run ./cmd/usersjson/main.go > ./data/users.json
	@echo "Done"