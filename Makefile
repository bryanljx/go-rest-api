## Variables
APP_BINARY_PATH ?= cmd/server/
APP ?= go-rest-api
DATABASE ?= postgres

DBHOST ?= 127.0.0.1
DBPORT ?= 5433
DBNAME ?= go-rest-api-db
DBUSER ?= dbuser
DBPASSWORD ?= 12345
DBSSLMODE ?= disable
# For dev, dont set ssl mode
DB_DSN ?= ${DATABASE}://${DBUSER}:${DBPASSWORD}@${DBHOST}:${DBPORT}/${DBNAME}?sslmode=${DBSSLMODE}

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# DB Tasks

.PHONY: initDB
initDB:
	docker run --name ${APP}-db -e POSTGRES_DB=${DBNAME} -e POSTGRES_USER=${DBUSER} -e POSTGRES_PASSWORD=${DBPASSWORD} -v pgdata:/var/lib/postgresql/data -p ${DBHOST}:${DBPORT}:5432 -d postgres:16.2
	sleep 3 # Let postgres start up first before continuing with setup

.PHONY: migrateDB
migrateDB:
	migrate -path="./database/migrations/" -database="${DB_DSN}" up

.PHONY: seedDB
seedDB:
	psql -h ${DBHOST} -p ${DBPORT} -U ${DBUSER} -d ${APP}-db -f ./database/seeds/init.sql

.PHONY: startDB
startDB:
	docker start ${APP}-db

.PHONY: deleteDB
deleteDB:
	-docker stop ${APP}-db
	docker rm ${APP}-db
	docker volume rm pgdata

.PHONY: resetDB
resetDB: deleteDB setupDB

## createMigrations name=$1: create a new database migration
.PHONY: createMigrations
createMigrations:
	@echo 'Creating migration files for ${name}...'
	migrate create -ext=.sql -dir=./database/migrations ${name}

## createMigrations name=$1: create a new database migration
.PHONY: createSeeds
createSeeds:
	@echo 'Creating migration files for ${name}...'
	migrate create -ext=.sql -dir=./database/seeds ${name}

# Development Tasks

.PHONY: build
build:
	go build -o=./tmp/bin/${BINARY_NAME} ./${APP_BINARY_PATH}

.PHONY: run
run:
	go run github.com/cosmtrek/air@v1.43.0 \
        --build.cmd "make build" --build.bin "./tmp/bin/server" --build.delay "100" \
        --build.exclude_dir "" \
        --build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
        --misc.clean_on_exit "true"

.PHONY: test
test:
	go test -v ./...

.PHONY: coverage
coverage:
	go test -v -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# Helper Tasks

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

# ## run/api: run the cmd/api application
# .PHONY: run/api
# run/api:
# 	go run ./cmd/api

# ## db/psql: connect to the database using psql
# .PHONY: db/psql
# db/psql:
# 	psql ${GREENLIGHT_DB_DSN}

# ## db/migrations/new name=$1: create a new database migration
# .PHONY: db/migrations/new
# db/migrations/new:
# 	@echo 'Creating migration files for ${name}...'
# 	migrate create -seq -ext=.sql -dir=./migrations ${name}

# ## db/migrations/up: apply all up database migrations
# .PHONY: db/migrations/up
# db/migrations/up: confirm
# 	@echo 'Running up migrations...'
# 	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

# # Change these variables as necessary.
# MAIN_PACKAGE_PATH := ./cmd/server
# BINARY_NAME := go-rest-api

# # ==================================================================================== #
# # QUALITY CONTROL
# # ==================================================================================== #

# ## tidy: format code and tidy modfile
# .PHONY: tidy
# tidy:
#     go fmt ./...
#     go mod tidy -v

# ## audit: run quality control checks
# .PHONY: audit
# audit:
#     go mod verify
#     go vet ./...
#     go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
#     go run golang.org/x/vuln/cmd/govulncheck@latest ./...
#     go test -race -buildvcs -vet=off ./...


# # ==================================================================================== #
# # DEVELOPMENT
# # ==================================================================================== #

# ## test: run all tests
# .PHONY: test
# test:
#     go test -v -race -buildvcs ./...

# ## test/cover: run all tests and display coverage
# .PHONY: test/cover
# test/cover:
#     go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
#     go tool cover -html=/tmp/coverage.out

# ## build: build the application
# .PHONY: build
# build:
#     # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
#     go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

# ## run: run the  application
# .PHONY: run
# run: build
#     /tmp/bin/${BINARY_NAME}

# ## run/live: run the application with reloading on file changes
# .PHONY: run/live
# run/live:
#     go run github.com/cosmtrek/air@v1.43.0 \
#         --build.cmd "make build" --build.bin "/tmp/bin/${BINARY_NAME}" --build.delay "100" \
#         --build.exclude_dir "" \
#         --build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
#         --misc.clean_on_exit "true"


# # ==================================================================================== #
# # OPERATIONS
# # ==================================================================================== #

# ## push: push changes to the remote Git repository
# .PHONY: push
# push: tidy audit no-dirty
#     git push

# ## production/deploy: deploy the application to production
# .PHONY: production/deploy
# production/deploy: confirm tidy audit no-dirty
#     GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=/tmp/bin/linux_amd64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
#     upx -5 /tmp/bin/linux_amd64/${BINARY_NAME}
#     # Include additional deployment steps here...
