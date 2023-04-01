.PHONY: gen-userlist parse-memp

gen-userlist:
	@echo "Generating userlist.json"
	@go run ./cmd/usersjson/main.go > ./data/userlist.json
	@echo "Done"

parse-memp:
	@echo "Parsing and Profiling memory"
	@go run ./cmd/parsetest/main.go -memprofile mem.prof
	@echo "Done"