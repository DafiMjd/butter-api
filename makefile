# Usage: make create_migration f=<file_name>
.PHONY: create_migration
create_migration:
ifneq ($(f),)
	@migrate create -ext sql -dir database/migrations $(f)
else
	$(error "Usage: make create_migration f=<file_name>")
endif

.PHONY: migration_up
migration_up: 
	@migrate -database "mysql://root:password@tcp(localhost:3306)/butter" -path database/migrations up

.PHONY: migration_down
migration_down: 
	@migrate -database "mysql://root:password@tcp(localhost:3306)/butter" -path database/migrations down

.PHONY: migration_force
migration_down: 
ifneq ($(f),)
	@migrate -database "mysql://root:password@tcp(localhost:3306)/butter" -path database/migrations force $(f)
else
	$(error "Usage: make migration_force f=<version>")
endif