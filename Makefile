build:
	go build -o bin/api.exe ./cmd/api/

run:
	go run ./cmd/api/

clean:
	rm -rf bin/
