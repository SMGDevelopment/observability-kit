NAME := rmp-observability-kit
TESTS=./...

test:
	go test $(TESTS) -coverprofile=c.out
	sed -i -e 's#^github.com/cheddartv/${NAME}/##g' c.out
