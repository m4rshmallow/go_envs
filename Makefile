test_clean:
	go test ./tests/...

test_dev:
	go test -cover -coverprofile=cover.out -coverpkg=./... ./tests/...

test_cover:
	go tool cover -html=cover.out