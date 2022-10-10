package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func GetIngredientByName(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		ingredient, _ := models.GetIngredientByName(name)
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
	// ingredient := make([]map[string]string, 0)
	ingredientRequest, _ := utils.ParseBodyToMap(request)

	if len(ingredientRequest["name"]) < 1 {
		writer.Write([]byte("Name cannot be empty."))
		writer.WriteHeader((http.StatusBadRequest))
		return
	}
	if len(ingredientRequest["skinTypes"]) > 0 {
		skinTypesRequest, _ := utils.ParseStringToArray(string(ingredientRequest["skinTypes"]))
		skinTypesArray := make([]models.SkinType, len(skinTypesRequest))
		fmt.Println(len(skinTypesRequest))
		for index, skinType := range skinTypesRequest {
			fmt.Println(string(skinType))
			st, _ := models.GetSkinTypeByName(string(skinType))
			skinTypesArray[index] = *st
			fmt.Println(index, skinType)
		}
		ingredient := &models.Ingredient{}
		ingredient.Name = utils.TrimQuotes(string(ingredientRequest["name"]))
		ingredient.Effect = utils.TrimQuotes(string(ingredientRequest["effects"]))
		ingredient.SkinTypes = skinTypesArray
		ingredient.Slug = utils.MakeSlug(ingredient.Name)

		str, _ := json.Marshal(ingredient)
		fmt.Println(string(str))

		ingr, err := ingredient.CreateIngredient()
		if ingr.ID > 0 {
			res, _ := json.Marshal(ingr)

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
			if IngredientToUpdate.Effect != "" {
				ingr.Effect = IngredientToUpdate.Effect
			}
			// if IngredientToUpdate.ForSkinTypes != nil {
			// ingr.ForSkinTypes = IngredientToUpdate.ForSkinTypes
			// }

			savedIngr, err := ingr.UpdateIngredient()
			if err == nil {
				res, _ := json.Marshal(savedIngr)

				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusOK)
				writer.Write(res)
			} else {
				if strings.Contains(err.Error(), "1062") {
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write([]byte("Ingredient with such name already exists."))
				} else {
					writer.WriteHeader(http.StatusInternalServerError)
				}
			}
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

		db := models.DeleteIngredient(deleteId)
		if db.Error != nil {
			fmt.Println(db.Error.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
	}
}
