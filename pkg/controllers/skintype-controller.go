package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nic_pan/skincare-ingredients/pkg/models"
)

func GetAllSkinTypes(writer http.ResponseWriter, request *http.Request) {
	// var skinTypes = []models.SkinType{
	// 	{ID: 1, Name: "Dry"},
	// 	{ID: 2, Name: "Oily"},
	// 	{ID: 3, Name: "Combination"},
	// 	{ID: 4, Name: "Dehydrated"},
	// 	{ID: 5, Name: "Sensitive"},
	// 	{ID: 6, Name: "Acne-Prone"},
	// 	{ID: 7, Name: "Normal"},
	// }
	skinTypes := models.GetAllSkinTypes()
	response, _ := json.Marshal(skinTypes)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetSkinType(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	dbId, _ := strconv.Atoi(id)
	skinType, _ := models.GetSkinTypeById(int64(dbId))

	response, _ := json.Marshal(skinType)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func AddSkinType(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// request.Body.Read()
	json.NewEncoder(writer).Encode("Created new skin type")
}
func UpdateSkinType(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	// TODO check if exists? and update in DB
	json.NewEncoder(writer).Encode("Updated " + id)
}
func DeleteSkinType(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	dbId, _ := strconv.Atoi(id)

	models.DeleteSkinType(int64(dbId))
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// writer.Write(response)
}
