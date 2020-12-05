build: fmt
	CGO_ENABLED=0 go build -o bin/ga -ldflags "-w -s"

copy: build
	sudo cp bin/ga /usr/local/bin/ga

fmt:
	go fmt ./...
