package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/controllers"
	"github.com/rs/zerolog/log"
)

var RegisterSkinTypeRoutes = func(router *mux.Router) {

	router.HandleFunc("/skinTypes", controllers.GetAllSkinTypes).Methods(http.MethodGet)
	router.HandleFunc("/skinType/{id}", controllers.GetSkinType).Methods(http.MethodGet)
	router.HandleFunc("/skinType", controllers.AddSkinType).Methods(http.MethodPost)
	router.HandleFunc("/skinType/{id}", controllers.UpdateSkinType).Methods(http.MethodPut)
	router.HandleFunc("/skinType/{id}", controllers.DeleteSkinType).Methods(http.MethodDelete)
	
	log.Debug().Msg("Registered Skin Type routes.")
}
