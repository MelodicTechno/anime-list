build:
	go build -o bin/anime-list-server.exe ./cmd/api/

run:
	go run ./cmd/api/

clean:
	rm -rf bin/
