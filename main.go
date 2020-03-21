package main

import (
	"encoding/json"
	"fmt"
	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/lruggieri/graphqltest/pkg/database"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "Nothing to be done here")
	})
	router.Handle("/graphiql", graphiqlHandler)
	router.Handle("/graphql", gqlHandler())

	server := &http.Server{
		Handler:      router,
		Addr:         ":3000",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Starting server...")
	log.Fatal(server.ListenAndServe())
}

// gqlHandler is the entry point for /graphql queries
func gqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", http.StatusBadRequest)
			return
		}

		var rBody struct {
			Query string `json:"query"`
		}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "%s", processQuery(rBody.Query))

	})
}

// processQuery processes graphQL queries
func processQuery(query string) (result string) {

	params := graphql.Params{
		Schema:        gqlSchema(),
		RequestString: query,
	}
	graphqlResult := graphql.Do(params)
	if len(graphqlResult.Errors) > 0 {
		fmt.Printf("failed to execute graphql operation, errors: %+v", graphqlResult.Errors)
	}
	rJSON, _ := json.Marshal(graphqlResult)

	return fmt.Sprintf("%s", rJSON)

}

// gqlSchema defines the schema for graphQL
func gqlSchema() graphql.Schema {
	fields := graphql.Fields{
		"jobs": &graphql.Field{
			Type:        graphql.NewList(jobType),
			Description: "All Jobs",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return database.GetAllJobs(), nil
			},
		},
		"job": &graphql.Field{
			Type:        jobType,
			Description: "Get Jobs by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return database.GetJobByID(params.Args["id"].(int)), nil
			},
		},
		"companies": &graphql.Field{
			Type:        graphql.NewList(companyType),
			Description: "All companies",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return database.GetAllCompanies(), nil
			},
		},
		"company": &graphql.Field{
			Type:        companyType,
			Description: "Get Company by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return database.GetCompanyByID(params.Args["id"].(int)), nil
			},
		},
		"people": &graphql.Field{
			Type:        graphql.NewList(personType),
			Description: "All people",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return database.GetAllPeople(), nil
			},
		},
		"person": &graphql.Field{
			Type:        personType,
			Description: "Get Person by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				return database.GetPersonByID(p.Args["id"].(int)), nil
			},
		},
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name:   "RootQuery",
				Fields: fields,
			},
		),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Printf("failed to create new schema, error: %v", err)
	}

	return schema

}
