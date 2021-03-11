name-generator: main.go bindata.go
	go build

bindata.go: data/*
	(cd data && go-bindata -pkg data *.txt)

run: bindata.go
	go run *.go

clean:
	rm data/bindata.go name-generator 2> /dev/null || true