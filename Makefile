run:
	#go run cmd/web/* -addr=":8081"
	go run `ls cmd/api/main.go | grep -v _test.go` -addr=":8081"

test:
	go test -v ./cmd/web/