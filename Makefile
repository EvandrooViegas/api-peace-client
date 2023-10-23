up:
	docker compose down && docker compose up -d 
down:
	docker compose down 
	
#app 

app-build: web-build api-build
app: net app-stop mongo mongoex web api
app-stop: mongo-stop mongoex-stop web-stop api-stop

#network 
net: 
	@echo --CREATING APP NETWORK ONEAPI-NETWORK
	-docker network create oneapi-network

# web
web-local:
	@echo --RUNNING WEB LOCALLY
	cd src/web && npm run dev
web: net web-stop
	@echo --STARTING WEB
	cd src/web && docker run --network oneapi-network --network-alias web --name oneapi-web -v .:/app -w /app -dp  3000:3000 oneapi-web sh -c "npm i && npm run dev"
web-stop:
	@echo --STOPPING WEB 
	-docker rm -f oneapi-web 
web-build: 
	@echo --BUILDING API
	cd src/web && docker build -t oneapi-web .


#mongo express
mongoex: mongoex-stop
	@echo --STARTING MONGO EXPRESS
	docker run -it --network oneapi-network \
	--network-alias mongoex \
    --name oneapi-mongoex \
    -dp 8081:8081 \
    -e ME_CONFIG_MONGODB_AUTH_DATABASE="oneapi" \
    -e ME_CONFIG_BASICAUTH_USERNAME="root" \
    -e ME_CONFIG_BASICAUTH_PASSWORD="rootpass" \
    mongo-express

mongoex-stop:
	@echo --STOPPING MONGO EXPRESS
	-docker stop oneapi-mongoex
	-docker remove oneapi-mongoex

# mongodb
mongo-stop:
	@echo --STOPPPING MONGODB
	-docker stop oneapi-mongodb 
	-docker remove oneapi-mongodb
mongo: net mongo-stop
	@echo --STARTING MONGODB
	-docker volume create oneapi-mongodb
	docker run --network oneapi-network --network-alias mongodb --name oneapi-mongodb -w /app --mount type=volume,src=oneapi-mongodb,target=/data/db -dp 27017:27017 mongo:latest 

# api
api-local: 
	@echo --RUNNING API LOCALLY API
	cd src/api/cmd/api/v1 && set APP_ENV=dev && air
api: net api-stop
	@echo --STARTING API
	cd src/api/cmd/api/v1  &&  docker run --network oneapi-network --network-alias api --name oneapi-api -dp 8080:8080 -w /app -v .:/app -e APP_ENV=dev oneapi-api
api-stop:
	@echo --STOPPING API
	-docker rm -f oneapi-api 

api-build: 
	@echo --BUILDING API
	cd src/api/cmd/api/v1 && docker build -t oneapi-api .
