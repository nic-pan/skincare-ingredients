package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/models"
	"github.com/nic_pan/skincare-ingredients/pkg/utils"
)

func GetAllIngredients(writer http.ResponseWriter, request *http.Request) {
	ingredients := models.GetAllIngredients()
	response, _ := json.Marshal(ingredients)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetIngredient(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		dbId, _ := utils.ParseID(id)
		ingredient, _ := models.GetIngredientById(dbId)
		if ingredient == nil {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			response, _ := json.Marshal(ingredient)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(response)
		}
	}
}

func AddIngredient(writer http.ResponseWriter, request *http.Request) {
	ingredient := &models.Ingredient{}
	utils.ParseBody(request, ingredient)
	ingr := ingredient.CreateIngredient()
	res, _ := json.Marshal(ingr)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func UpdateIngredient(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		updateId, _ := utils.ParseID(id)

		IngredientToUpdate := &models.Ingredient{}
		utils.ParseBody(request, IngredientToUpdate)

		ingr, _ := models.GetIngredientById(updateId)
		if ingr != nil {
			if IngredientToUpdate.Name != "" {
				ingr.Name = IngredientToUpdate.Name
			}

			res, _ := json.Marshal(ingr)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(res)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}

func DeleteIngredient(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		deleteId, _ := utils.ParseID(id)

		models.DeleteIngredient(deleteId)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
	}
}
