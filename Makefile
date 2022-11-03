name-generator: main.go
	go build

run:
	go run *.go

clean:
	rm name-generator 2> /dev/null || true