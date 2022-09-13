package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/routes"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Debug().Msg("Starting server...")
	startServer()
	// config.ConnectToDB()
}

func startServer() {

	router := mux.NewRouter().StrictSlash(true)

	log.Debug().Msg("Registering routes.")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome, skincare\n")
	})
	routes.RegisterSkinTypeRoutes(router)
	routes.RegisterIngredientRoutes(router)
	routes.RegisterCombinationRoutes(router)

	log.Fatal().Err(http.ListenAndServe(":8081", router))
	log.Debug().Msg("Server started.")
}
