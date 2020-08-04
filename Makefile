VERSION :=v0.0.4

.PHONY: build install release release-linux release-mac


build:
	go build -ldflags '-w -s' -o file_transfer

install: build
	cp -f file_transfer /usr/local/bin/


release: release-linux release-mac

release-linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o file_transfer
	tar zcvf file_transfer-$(VERSION)-linux-amd64.tar.gz ./file_transfer
	rm file_transfer


release-mac:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o file_transfer
	tar zcvf file_transfer-$(VERSION)-darwin-amd64.tar.gz ./file_transfer
	rm file_transfer






