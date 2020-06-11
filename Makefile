.PHONY: cmd clean

cmd: $(wildcard ./cmd/*.go ./core/*.go ./version/*.go ./*.go)
	go build -o scanner;
	./scanner -h

clean:
	rm scanner