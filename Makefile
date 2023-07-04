POSTGRES_HOST=localhost
POSTGRES_USER=l0_back_user
POSTGRES_PASSWORD=l0_back_hard_password
POSTGRES_DB=l0_back_db
DBSTRING="host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable"
DRIVER=postgres

# make mig_create name=YOUR_MIGRATION_NAME
mig_create:
	goose -dir ./migrations $(DRIVER) $(DBSTRING)  create $(name) sql

mig_status:
	goose -dir ./migrations $(DRIVER) $(DBSTRING)  status

pub:
	go run ~/go/pkg/mod/github.com/nats-io/stan.go@v0.10.4/examples/stan-pub/main.go \
	-s localhost:4223 -c cluster -id 1  models "$$(cat /home/dimyasha/GolandProjects/l0_wb/model/example.json)"