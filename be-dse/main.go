package main

import (
	"be-dse/middleware"
	"be-dse/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	r := router.Router()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handler := cors.Handler(middleware.NewAuthMiddleware(r))

	fmt.Println("Server dijalankan pada port 3000...")
	log.Fatal(http.ListenAndServe(":3000", handler))

}
