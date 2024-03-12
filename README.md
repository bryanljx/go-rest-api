# go-rest-api

This is a simple rest api implemented in go for Govtech OneCV take home assignment.

## Prerequisites
- go v1.22
- docker
- psql
- make
- go-migrate

## Set up dev environment
To simplify setting up a local dev environment, run the following commands when starting the project for the first time:

1. Create a docker container for postgres database

```sh
make initDB
```

2. Run migrations with `go-migrate`

```sh
make migrateDB
```

3. Seed db 

```sh
make seedDB
```

To destroy the environment, simply run:

```sh
make deleteDB
```

This will delete the docker container used for the db.

NOTE: The set up with this particular Makefile has not been tested thoroughly. In particular, support for other platforms/os has not been verified. This was tested using a machine running Fedora 39.