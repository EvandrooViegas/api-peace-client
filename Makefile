.PHONY: client server

test-client:
	cd ./client && clear && npm run test

up:
	docker compose up -d
down:
	docker compose down 
client:
	cd ./client && npm run dev

server:
	 cd src/api/cmd/api/v1 && nodemon --exec go run main.go