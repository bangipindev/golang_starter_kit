DB_URL=mysql://wsluser:wslpassword@tcp(gpt_mysql:3306)/golang?multiStatements=true
MIGRATE=migrate -path migrations -database "$(DB_URL)"

migrate-version:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

# migrate-reset:
# 	$(MIGRATE) drop -f && $(MIGRATE) up

# migrate-force:
# 	$(MIGRATE) force 1
