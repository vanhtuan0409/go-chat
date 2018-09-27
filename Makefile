build:
	GOOS=darwin GOARCH=amd64 go build -o bin/gchat-server github.com/vanhtuan0409/go-chat/cmd/server
	GOOS=darwin GOARCH=amd64 go build -o bin/gchat-client github.com/vanhtuan0409/go-chat/cmd/client

release:
	make build
	tar -C bin -zcvf dist/gchat-server-macos.tar.gz gchat-server
	tar -C bin -zcvf dist/gchat-client-macos.tar.gz gchat-client
