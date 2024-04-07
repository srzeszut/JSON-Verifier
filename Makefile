

all: test verify

test:
	go test ./src

verify: src/verify.go src/main.go
	go build -o verify ./src
	./verify $(FILE)

clean:
	rm -f verify


