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
<<<<<<< HEAD
=======
	resolvers "github.com/Caicrs/Payfood-backend/graph/resolvers"
>>>>>>> main
	"github.com/joho/godotenv"
)

const defaultPort = "4000"

func main() {
	godotenv.Load()
	myDb := os.Getenv("POSTGRES_URL")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

<<<<<<< HEAD
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{}))
=======
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
>>>>>>> main

	customCtx := &common.CustomContext{
		Database: db,
	}

	fmt.Println(myDb)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
