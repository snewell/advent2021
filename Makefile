.PHONY := help format

help:
	@echo "Usage: make <day>[a|b]"
	@echo "example: make 1a"

%: %.go
	go run $^ input/$(shell echo $^ | grep -o '^[[:digit:]]*')

format:
	find -name '*.go' | xargs gofmt -w
