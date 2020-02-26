# Golang API

## Summary

* This is a microservice example (1 endpoint) wich uses a clean architecture
  approach by separating concerns between model, repository and business logic. 

* The testing approach is TDT (Table Driven Tests)

* There is an interface for the repository in the v1/main.go file so that the
  persistence can switch between technologies wihtout affecting the business
  logic

* It also uses the
  [https://github.com/joho/godotenv](https://github.com/joho/godotenv) package
  for support to .env files

## Future features

* Support for mongoDB as persistence technology
* Dockerization

