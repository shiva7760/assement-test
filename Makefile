all:
	docker compose up -d && go build . && ./assement-test