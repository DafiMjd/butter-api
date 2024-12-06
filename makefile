# Usage: make create_migration_mysql f=<file_name>
.PHONY: create_migration_mysql
create_migration_mysql:
ifneq ($(f),)
	@migrate create -ext sql -dir database/migrations/mysql $(f)
else
	$(error "Usage: make create_migration_mysql f=<file_name>")
endif

# Usage: make create_migration_postgres f=<file_name>
.PHONY: create_migration_postgres
create_migration_postgres:
ifneq ($(f),)
	@migrate create -ext sql -dir database/migrations/postgres -seq $(f)
else
	$(error "Usage: make create_migration_postgres f=<file_name>")
endif

.PHONY: migration_up_mysql
migration_up_mysql:
	@migrate -database "mysql://root:password@tcp(localhost:3306)/butter" -path database/migrations/mysql up

.PHONY: migration_down_mysql
migration_down_mysql: 
	@migrate -database "mysql://root:password@tcp(localhost:3306)/butter" -path database/migrations/mysql down

.PHONY: migration_force_mysql
migration_down_mysql: 
ifneq ($(f),)
	@migrate -database "mysql://root:password@tcp(localhost:3306)/butter" -path database/migrations/mysql force $(f)
else
	$(error "Usage: make migration_force_mysql f=<version>")
endif


.PHONY: postgres_migration_up
postgres_migration_up:
	@migrate -database "postgres://postgres:$(p)@localhost:5432/butter?sslmode=disable" -path database/migrations/postgres up

.PHONY: postgres_migration_down
postgres_migration_down: 
	@migrate -database "postgres://postgres:$(p)@localhost:5432/butter?sslmode=disable" -path database/migrations/postgres down

.PHONY: postgres_migration_force
postgres_migration_down: 
ifneq ($(f), $(p),)
	@migrate -database "postgres://postgres:$(p)@localhost:5432/butter?sslmode=disable" -path database/migrations/postgres force $(f)
else
	$(error "Usage: make postgres_migration_force f=<version> p=<password>")
endif