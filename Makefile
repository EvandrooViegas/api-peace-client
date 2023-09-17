.PHONY: client server

test-client:
	cd ./client && clear && npm run test

up:
	docker compose up -d
down:
	docker compose down 
web:
	cd src/web && npm run dev

api:
	 cd src/api/cmd/api/v1 && nodemon --exec go run main.go