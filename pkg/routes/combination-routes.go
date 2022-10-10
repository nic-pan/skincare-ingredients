package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/controllers"
	"github.com/rs/zerolog/log"
)

var RegisterCombinationRoutes = func(router *mux.Router) {

	router.HandleFunc("/combinations", controllers.GetAllCombinations).Methods(http.MethodGet)
	router.HandleFunc("/combination/{id}", controllers.GetCombination).Methods(http.MethodGet)
	router.HandleFunc("/combination", controllers.GetCombinationOfIngredients).Methods(http.MethodGet)
	router.HandleFunc("/combination", controllers.AddCombination).Methods(http.MethodPost)
	// router.HandleFunc("/combination/{id}", controllers.UpdateIngredient).Methods(http.MethodPut)
	router.HandleFunc("/combination/{id}", controllers.DeleteCombination).Methods(http.MethodDelete)

	log.Debug().Msg("Registered Combination routes.")
}
