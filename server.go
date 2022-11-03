package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Caicrs/Payfood-backend/common"
	"github.com/Caicrs/Payfood-backend/graph/generated"
	graph "github.com/Caicrs/Payfood-backend/graph/resolvers"
	"github.com/joho/godotenv"
)

const defaultPort = "4000"

func main() {

	godotenv.Load()
	myDb := os.Getenv("POSTGRES_URL")
	fmt.Println(myDb)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	customCtx := &common.CustomContext{
		Database: db,
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
