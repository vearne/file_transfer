build:
	go build -ldflags '-w -s' -o file_transfer

install: build
	cp -f file_transfer /usr/local/bin/
