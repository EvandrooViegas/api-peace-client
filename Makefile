up:
	docker compose down && docker compose up -d 
down:
	docker compose down 
	
#app 

app-build: web-build api-build
app-run: web-start

# web
web-start: 
	@echo "Starting oneapi-web"
	cd src/web && docker run --name oneapi-web -v .:/app -w /app -d -p  3000:3000 -p 3001:3001 oneapi-web sh -c "npm run dev"
web-stop:
	@echo "Stopping and removing oneapi-web..."; 
	docker rm -f oneapi-web 
web-build: 
	cd src/web && docker build -t oneapi-web .
# api
api-local:
	cd src/api/cmd/api/v1 && npm i -g nodemon && go mod tidy && set APP_ENV=dev&& nodemon --exec go run main.go 
api-run: 
	cd src/api/cmd/api/v1 && docker run --name oneapi-api --network oneapi-network --network-alias api -dp 8080:8080 -w /app --mount "type=bind,src=.,target=/app" golang:1.21.1 sh -c "go mod tidy && go run main.go" 
api-build: 
	cd src/api/cmd/api/v1 && docker build -t oneapi-api .
