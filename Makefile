run:
	#go run cmd/web/* -addr=":8081"
	go run `ls cmd/api/*.go | grep -v _test.go`

test:
	go test -v ./cmd/web/