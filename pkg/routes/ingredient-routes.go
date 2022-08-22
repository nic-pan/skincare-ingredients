package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/controllers"
	"github.com/rs/zerolog/log"
)

var RegisterIngredientRoutes = func(router *mux.Router) {

	router.HandleFunc("/ingredients", controllers.GetAllIngredients).Methods(http.MethodGet)
	router.HandleFunc("/ingredient/{id}", controllers.GetIngredient).Methods(http.MethodGet)
	router.HandleFunc("/ingredient", controllers.AddIngredient).Methods(http.MethodPost)
	router.HandleFunc("/ingredient/{id}", controllers.UpdateIngredient).Methods(http.MethodPut)
	router.HandleFunc("/ingredient/{id}", controllers.DeleteIngredient).Methods(http.MethodDelete)

	log.Debug().Msg("Registered Ingredient routes.")
}
