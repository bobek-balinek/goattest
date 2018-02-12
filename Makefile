build:
	go build -o dist/goattest ./cmd/goattest

clean:
	rm -fr dist/

.PHONY: build clean
