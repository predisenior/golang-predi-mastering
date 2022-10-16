package main

import (
	"log"
	"net/http"
	"os"
	"project-3/config"
	"project-3/graph"
	"project-3/graph/generated"
	"project-3/routes"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
    config.KonekDB()
    // menjalankan fungsi restapi
    go func(){
          e:=routes.New()
    // restapi:=e.Start(":3000")
    http.ListenAndServe("localhost:3000",e)
    }()


    // graphql
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))





}
