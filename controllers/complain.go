package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"safety/models"
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
