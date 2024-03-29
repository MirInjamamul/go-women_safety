package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"safety/models"
	"safety/utils"

	"github.com/gorilla/mux"
)

type ComplainInput struct {
	Thana        string `json:"thana"`
	ComplainType string `json:"complain_type"`
	Comment      string `json:"comment"`
}

func CreateComplain(w http.ResponseWriter, r *http.Request) {
	var input ComplainInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	complain := &models.Complain{
		Thana:        input.Thana,
		ComplainType: input.ComplainType,
		Comment:      input.Comment,
	}

	models.DB.Create(complain)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(complain)
}

func GetAllComplains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var complains []models.Complain
	models.DB.Find(&complains)

	json.NewEncoder(w).Encode(complains)
}

func GetComplain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var complain models.Complain

	if err := models.DB.Where("id = ?", id).First(&complain).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Complain Not Found")
		return
	}

	json.NewEncoder(w).Encode(complain)
}

func DeleteComplain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var complain models.Complain

	if err := models.DB.Where("id = ?", id).First(&complain).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	models.DB.Delete(&complain)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(complain)
}
