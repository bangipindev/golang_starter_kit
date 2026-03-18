DB_URL=mysql://root:root@tcp(localhost:3306)/golang?multiStatements=true
MIGRATE=migrate -path migrations -database "$(DB_URL)"

migrate-version:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-reset:
	$(MIGRATE) drop -f && $(MIGRATE) up

migrate-force:
	$(MIGRATE) force 1
