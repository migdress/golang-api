# Golang API

## Summary

* This is a golang api example (1 endpoint) wich uses a clean architecture
  approach by separating concerns between model, repository and business logic. 

* The testing approach is TDT (Table Driven Tests)

* There is an interface for the repository in the v1/main.go file so that the
  persistence can switch between technologies without affecting the business
  logic

* It also uses the
  [https://github.com/joho/godotenv](https://github.com/joho/godotenv) package
  for support to .env files

* There is a working mongodb repository implementation, tests are missing
  though. The host, port, database and collection can be configured in the `.env` file

## How to test it

* Clone the repository with `git clone https://github.com/migdress/golang-api.git`
* Cd into directory `cd golang-api`
* Create the .env file `cp .env.example .env`
* Configure env vars for mongodb connection in `.env`
    * **NOTE**: It's necessary (for now) to run a separate mongodb server and configure env vars in the `.env` to connect to it
* Test it with `make run`
* Run it in a container with `make dockerrun`
* Try making POST requests with a REST client to `localhost:8080`

## Future features

* Integration tests with mongodb
* Dockerization with docker-compose including the mongodb instance, for now there is a Dockerfile for the golang-api



