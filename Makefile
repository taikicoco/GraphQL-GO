build:
	docker-compose build
build-nc:
	docker-compose build --no-cache
up:
	docker-compose up
down:
	docker-compose down

sh:
	docker-compose exec backend sh
mysql:
	docker compose exec mysql mysql -uroot -ppasswordroot -Ddb