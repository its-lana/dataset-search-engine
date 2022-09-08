package main

import (
	"be-dse/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := router.Router()

	fmt.Println("Server dijalankan pada port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))

}
