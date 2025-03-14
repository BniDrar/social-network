
# run:
# 	export ADDR=":9999" && go run ./cmd/web/ -addr=$$ADDR
test:
	### you can see exactly which tests are being run by using the -v flag to
	### get the verbose output:
	# go test -v 
	go test -v -failfast=false

test-all:
	go test -race 
test-perf:
	go test -count=100 
push:
	git push github
	git push origin
config:
	go mod tidy
	# mkdir tls
	# go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost --outdir=./tls