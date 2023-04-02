.PHONY: gen-userlist memp1 memp2 memp3 memp4 memp5

gen-userlist:
	@echo "Generating userlist.json"
	@go run ./cmd/usersjson/main.go > ./data/userlist.json
	@echo "Done"

memp1:
	./memprof.sh 1

memp2:
	./memprof.sh 2

memp3:
	./memprof.sh 3

memp4:
	./memprof.sh 4

memp5:
	./memprof.sh 5

memp_all: memp1 memp2 memp3 memp4 memp5
