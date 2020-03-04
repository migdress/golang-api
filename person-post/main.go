package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-api/person-post/model"
	"golang-api/person-post/repository"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/xeipuuv/gojsonschema"
)

type peopleRepositoryInterface interface {
	Save(m model.Person) error
}

type handler func(w http.ResponseWriter, r *http.Request)

type requestBody struct {
	FullName  *string `json:"full_name"`
	DNI       *string `json:"dni"`
	Birthdate *string `json:"birthdate"`
}

var validationSchema = gojsonschema.NewStringLoader(`
	{
		"type": "object",
		"required": [
			"full_name",	
			"dni",	
			"birthdate"
		],
		"properties": {
			"full_name": {
				"type": "string",
				"minLength": 1
			},
			"dni": {
				"type": "string",
				"minLength": 1
			},
			"birthdate": {
				"type": "string",
				"minLength": 1
			}
		}
	}
`)

func adapter(repo peopleRepositoryInterface) handler {
	return func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		var requestBody requestBody
		err := json.Unmarshal([]byte(body), &requestBody)
		if err != nil {
			writeResponse(
				http.StatusBadRequest,
				fmt.Sprintf(`{"error":"%v"}`, err),
				w,
			)
			return
		}

		dataToValidate := gojsonschema.NewGoLoader(requestBody)
		result, _ := gojsonschema.Validate(validationSchema, dataToValidate)
		if !result.Valid() {
			errorsSlice := []string{}
			for _, anError := range result.Errors() {
				errString := fmt.Sprintf("%v", anError)
				errorsSlice = append(errorsSlice, errString)
			}
			theErrors, _ := json.Marshal(map[string]interface{}{
				"errors": errorsSlice,
			})
			writeResponse(
				http.StatusBadRequest,
				string(theErrors),
				w,
			)
			return
		}

		person := model.Person{
			FullName:  *requestBody.FullName,
			DNI:       *requestBody.DNI,
			Birthdate: *requestBody.Birthdate,
		}

		err = repo.Save(person)
		if err != nil {
			writeResponse(
				http.StatusInternalServerError,
				fmt.Sprintf(`{"error":"%v"}`, err),
				w,
			)
			return
		}

		writeResponse(
			http.StatusOK,
			"",
			w,
		)
	}
}

func writeResponse(code int, msg string, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	host := os.Getenv("MONGODB_HOST")
	if strings.TrimSpace(host) == "" {
		panic("MONGODB_HOST is empty")
	}

	port := os.Getenv("MONGODB_PORT")
	if strings.TrimSpace(port) == "" {
		panic("MONGODB_PORT is empty")
	}

	database := os.Getenv("MONGODB_DATABASE")
	if strings.TrimSpace(database) == "" {
		panic("MONGODB_DATABASE is empty")
	}

	collection := os.Getenv("MONGODB_COLLECTION")
	if strings.TrimSpace(collection) == "" {
		panic("MONGODB_COLLECTION is empty")
	}

	serveInPort := os.Getenv("APP_PORT")
	if strings.TrimSpace(serveInPort) == "" {
		panic("APP_PORT is empty")
	}

	client, err := repository.NewMongoDBConnection(
		host,
		port,
	)
	if err != nil {
		panic(err)
	}

	repo := repository.NewPeopleRepository(
		client,
		database,
		collection,
		context.Background(),
	)

	http.HandleFunc("/", adapter(repo))
	log.Fatal(http.ListenAndServe(":"+serveInPort, nil))
}
