.PHONY: cmd clean

cmd: $(wildcard ./cmd/*.go ./core/*.go ./version/*.go ./*.go)
	go build -ldflags "-s -w" -o scanner;

clean:
	rm scanner