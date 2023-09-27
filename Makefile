NAME := rmp-observability-kit
TESTS=./...

test:
	go test $(TESTS) -coverprofile=c.out
	sed -i -e 's#^github.com/cheddartv/${NAME}/##g' c.out

test-html:
	go test $(TESTS) -tags=test -coverprofile=c.out
	go tool cover -html=c.out
