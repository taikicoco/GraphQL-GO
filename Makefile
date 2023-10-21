build:
	docker-compose build
build-nc:
	docker-compose build --no-cache
up:
	docker-compose up
upd:
	docker-compose up -d
down:
	docker-compose down
log:
	docker-compose logs -f
sh:
	docker-compose exec backend sh


# db
MYSQL_USER := user
MYSQL_PASSWORD := password
MYSQL_DB := db

mysql:
	docker compose exec mysql mysql -uroot -ppasswordroot -Ddb

migrate/new:
	echo '-- +goose Up' > db/migration/_.sql

migrate/up:
	docker-compose exec -T backend sh ../db/migration/script/migrate-up.sh 
