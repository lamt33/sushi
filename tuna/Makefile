include .env
export

compose_up:
	docker-compose --env-file .env --build

compose_up_d:
	docker-compose  --env-file .env -d --build

compose_down: 
	docker-compose down

remove_images:
	docker rm -vf $(docker ps -aq)
	docker rmi -f $(docker images -aq)
	rm -rf postgres-data
	mkdir postgres-data

dev:
	docker-compose up  --env-file .env -d database redis datadog
	go run ./tuna_api/cmd/app.go
