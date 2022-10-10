package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/models"
	"github.com/nic_pan/skincare-ingredients/pkg/utils"
)

func GetAllCombinations(writer http.ResponseWriter, request *http.Request) {
	combinations := models.GetAllCombinations()
	response, _ := json.Marshal((combinations))

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetCombination(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		dbId, _ := utils.ParseID(id)
		ingredient, _ := models.GetCombinationById(dbId)
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

func GetCombinationOfIngredients(writer http.ResponseWriter, request *http.Request) {
	ingredients := strings.Split(request.URL.Query().Get("ingredients[]"), ",")
	if len(ingredients) > 2 {
		writer.Write([]byte("Too many ingredients provided in request."))
		writer.WriteHeader(http.StatusBadRequest)
	} else if len(ingredients) < 2 {
		writer.Write([]byte("Not enough ingredients. Please provide 2 in your request."))
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		ingr1, _ := models.GetIngredientByName(string(ingredients[0]))
		ingr2, _ := models.GetIngredientByName(string(ingredients[1]))

		combination := models.GetCombinationForIngredients(ingr1, ingr2)
		if combination == nil {
			writer.Write([]byte("No information available for these ingredients."))
			writer.WriteHeader(http.StatusNotFound)
		} else {
			response, _ := json.Marshal(combination)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(response)
		}
	}
}

func AddCombination(writer http.ResponseWriter, request *http.Request) {
	combinationRequest, _ := utils.ParseBodyToMap(request)
	ingredients, _ := utils.ParseStringToArray(string(combinationRequest["ingredients"]))
	compatible, err := strconv.ParseBool(string(combinationRequest["isCompatible"]))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		reason := utils.TrimQuotes(string(combinationRequest["reason"]))
		fmt.Println(compatible)
		fmt.Println(reason)
		if len(ingredients) > 2 {
			writer.Write([]byte("Too many ingredients provided in request."))
			writer.WriteHeader(http.StatusBadRequest)
		} else if len(ingredients) < 2 {
			writer.Write([]byte("Not enough ingredients. Please provide 2 in your request."))
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			ingr1, _ := models.GetIngredientByName(string(ingredients[0]))
			ingr2, _ := models.GetIngredientByName(string(ingredients[1]))

			if ingr1 == nil || ingr2 == nil {
				writer.WriteHeader(http.StatusBadRequest)
			}

			combination := &models.Combination{Ingredient1: ingr1.ID, Ingredient2: ingr2.ID, IsCompatible: compatible, Reason: reason}
			result, err := combination.CreateCombination()

			if result.ID > 0 {
				res, _ := json.Marshal(result)

				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusOK)
				writer.Write(res)
			} else {
				if err != nil {
					if strings.Contains(err.Error(), "1062") {
						writer.WriteHeader(http.StatusBadRequest)
						writer.Write([]byte("Ingredient with such name already exists."))
					} else {
						writer.WriteHeader(http.StatusInternalServerError)
					}
				} else {
					writer.WriteHeader(http.StatusBadRequest)
				}
			}
		}
	}
}
