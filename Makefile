.PHONY: gen-userlist parse1-memp parse2-memp

gen-userlist:
	@echo "Generating userlist.json"
	@go run ./cmd/usersjson/main.go > ./data/userlist.json
	@echo "Done"

parse1-memp:
	@echo "Parsing(Case1) and Memory Profiling"
	@go run ./cmd/parsetest/main.go -case 1 -memprofile mem1.prof
	@go tool pprof -png mem1.prof > mem1.prof.png
	@echo "Done"

parse2-memp:
	@echo "Parsing(Case2) and Memory Profiling"
	@go run ./cmd/parsetest/main.go -case 2 -memprofile mem2.prof
	@go tool pprof -png mem2.prof > mem2.prof.png
	@echo "Done"
