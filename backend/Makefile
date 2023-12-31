# variables

bin = app 

build:
	go build -o ./bin/$(bin) ./cmd/main.go

clean:
	rm ./bin/*

brun: build
	./bin/$(bin)

run:
	./bin/$(bin)

dev:
	go run cmd/main.go
