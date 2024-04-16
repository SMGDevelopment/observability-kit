NAME := observability-kit
TESTS=./...

test:
	go test $(TESTS) -coverprofile=c.out
	sed -i -e 's#^github.com/SMGDevelopment/${NAME}/##g' c.out

test-html:
	go test $(TESTS) -tags=test -coverprofile=c.out
	go tool cover -html=c.out