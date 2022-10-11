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

func GetAllSkinTypes(writer http.ResponseWriter, request *http.Request) {
	skinTypes := models.GetAllSkinTypes()
	response, _ := json.Marshal(skinTypes)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetSkinType(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	dbId, _ := utils.ParseID(id)
	skinType, _ := models.GetSkinTypeById(dbId)
	if skinType == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	response, _ := json.Marshal(skinType)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func AddSkinType(writer http.ResponseWriter, request *http.Request) {
	SkinType := &models.SkinType{}
	utils.ParseBody(request, SkinType)

	if SkinType.Name == "" {
		writer.Write([]byte("Name cannot be empty."))
		writer.WriteHeader((http.StatusBadRequest))
		return
	}
	SkinType.Slug = utils.MakeSlug(SkinType.Name)

	st, err := SkinType.CreateSkinType()

	if st.ID > 0 {
		res, _ := json.Marshal(st)
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
	writer.Write([]byte("Skin Type with such name already exists."))
}

func UpdateSkinType(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updateId, _ := utils.ParseID(id)

	SkinTypeToUpdate := &models.SkinType{}
	utils.ParseBody(request, SkinTypeToUpdate)

	st, _ := models.GetSkinTypeById(updateId)
	if st == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if SkinTypeToUpdate.Name != "" {
		st.Name = SkinTypeToUpdate.Name
	}
	if SkinTypeToUpdate.Characteristics != "" {
		st.Characteristics = SkinTypeToUpdate.Characteristics
	}
	savedST, err := st.UpdateSkinType()
	if err == nil {
		res, _ := json.Marshal(savedST)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(res)
		return
	}
	if !strings.Contains(err.Error(), "1062") {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte("Skin Type with such name already exists."))
}

func DeleteSkinType(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	deleteId, _ := utils.ParseID(id)

	db := models.DeleteSkinType(deleteId)
	if db.Error != nil {
		fmt.Println(db.Error.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
