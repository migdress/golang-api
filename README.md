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
  though. The host, port, database and collection can be configured in .env file

## Future features

* Integration tests with mongodb
* Dockerization

