all: install

fmt:
	gofmt -l -w -s ./

clean:
	go clean -i ./
	rm -rf output

test:
    go test

install: clean fmt test
	echo "check finished"
