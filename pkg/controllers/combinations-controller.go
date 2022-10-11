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
	var resp []any
	for _, combination := range combinations {
		convertedCombination := convertCombinationResponse(&combination)
		resp = append(resp, convertedCombination)
	}
	response, _ := json.Marshal(resp)
	writer.Header().Set("Content-Type", "application/json")
	if string(response) == "null" {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetCombination(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	dbId, _ := utils.ParseID(id)
	combination, _ := models.GetCombinationById(dbId)
	if combination == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	resp := convertCombinationResponse(combination)
	response, _ := json.Marshal(resp)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetCombinationOfIngredients(writer http.ResponseWriter, request *http.Request) {
	ingredients := strings.Split(request.URL.Query().Get("ingredients[]"), ",")
	if len(ingredients) > 2 {
		writer.Write([]byte("Too many ingredients provided in request."))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(ingredients) < 2 {
		writer.Write([]byte("Not enough ingredients. Please provide 2 in your request."))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	ingr1, _ := models.GetIngredientByName(string(ingredients[0]))
	ingr2, _ := models.GetIngredientByName(string(ingredients[1]))

	combination := models.GetCombinationForIngredients(ingr1, ingr2)
	if combination == nil {
		writer.Write([]byte("No information available for these ingredients."))
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	resp := convertCombinationResponse(combination)
	response, _ := json.Marshal(resp)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func AddCombination(writer http.ResponseWriter, request *http.Request) {
	combinationRequest, _ := utils.ParseBodyToMap(request)
	ingredients, _ := utils.ParseStringToArray(string(combinationRequest["ingredients"]))
	compatible, err := strconv.ParseBool(string(combinationRequest["isCompatible"]))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	reason := utils.TrimQuotes(string(combinationRequest["reason"]))
	fmt.Println(compatible)
	fmt.Println(reason)
	if len(ingredients) > 2 {
		writer.Write([]byte("Too many ingredients provided in request."))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(ingredients) < 2 {
		writer.Write([]byte("Not enough ingredients. Please provide 2 in your request."))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
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
		return
	}
	if err == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if !strings.Contains(err.Error(), "1062") {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte("Ingredient with such name already exists."))
}

func UpdateCombination(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updateId, _ := utils.ParseID(id)
	combination, _ := models.GetCombinationById(updateId)
	if combination == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	updateRequest, _ := utils.ParseBodyToMap(request)
	ingredients, _ := utils.ParseStringToArray(string(updateRequest["ingredients"]))

	ingredient1, _ := models.GetIngredientByName(string(ingredients[0]))
	ingredient2, _ := models.GetIngredientByName(string(ingredients[1]))

	combination.Ingredient1 = ingredient1.ID
	combination.Ingredient2 = ingredient2.ID
	combination.Reason = utils.TrimQuotes(string(updateRequest["reason"]))
	compatible, _ := strconv.ParseBool(string(updateRequest["isCompatible"]))
	combination.IsCompatible = compatible

	savedCombination, err := combination.UpdateCombination()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(savedCombination)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteCombination(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	deleteId, _ := utils.ParseID(id)

	_, db := models.DeleteCombination(deleteId)
	if db.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func convertCombinationResponse(combination *models.Combination) map[string]any {
	var ingredients [2]string
	ingr1, _ := models.GetIngredientById(combination.Ingredient1)
	ingredients[0] = ingr1.Name
	ingr2, _ := models.GetIngredientById(combination.Ingredient2)
	ingredients[1] = ingr2.Name

	response := make(map[string]any)
	response["id"] = combination.ID
	response["ingredients"] = ingredients
	response["reason"] = combination.Reason
	response["isCompatible"] = combination.IsCompatible

	return response
}
