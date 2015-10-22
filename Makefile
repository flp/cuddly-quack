all:
	go fmt ./...
	go build ./...
	go build -o bin/cuddly-quack

clean:
	rm -rf bin
