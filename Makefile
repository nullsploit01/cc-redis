build-server:
	cd ./server && go build -o ccredis-server

build-cli:
	cd ./cli && go build -o ccredis-cli

run-server: build-server
	cd ./server && ./ccredis-server

run-cli: build-cli 
	cd ./cli && ./ccredis-cli
