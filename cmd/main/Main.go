package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/routes"
)

func main() {
	startServer()

}

func startServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome, skincare\n")
	})
	routes.RegisterSkinTypeRoutes(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
