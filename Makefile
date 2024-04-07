

all: test verify

test:
	go test ./src

build: src/verify.go src/main.go
	go build -o verify ./src
verify: build
	./verify $(FILE)

clean:
	rm -f verify


