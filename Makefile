name-generator: main.go bindata.go
	go build

bindata.go: data/*
	go-bindata data/

run: bindata.go
	go run *.go

clean:
	rm bindata.go
	rm name-generator