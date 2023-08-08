bin/dpsd-import:
	go get .
	go build -o bin/dpsd-import

clean:
	rm -rf bin

.PHONY: clean
