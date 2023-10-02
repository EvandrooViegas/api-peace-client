.PHONY: client server


# docker 
up:
	docker compose up -d
down:
	docker compose down 
d-restart: 
	docker compose down && docker compose up -d 

# web
web:
	cd src/web && set && npm run dev 
web-d-build: 
	cd src/web && docker build -t oneapi-web .

# api
api:
	cd src/api/cmd/api/v1 && npm i -g nodemon && go mod tidy && set APP_ENV=dev&& nodemon --exec go run main.go 
api-d-build: 
	cd src/api/cmd/api/v1 && docker build -t oneapi-api .