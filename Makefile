.PHONY: gen-userlist

gen-userlist:
	@echo "Generating userlist.json"
	@go run ./cmd/usersjson/main.go > ./data/userlist.json
	@echo "Done"