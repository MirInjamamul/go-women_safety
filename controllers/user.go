package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"safety/models"
	"safety/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	models.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := mux.Vars(r)["email"]
	var user models.User

	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User Not Found")
		return
	}

	json.NewEncoder(w).Encode(user)
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Mobile:   input.Mobile,
		Password: input.Password,
	}

	// Generate bearer Token for the user
	token, err := GenerateBearerToken(*user)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Token Generation Failed")
		return
	}

	//store the token
	user.Token = token

	models.DB.Create(user)

	w.Header().Set("Content-Type", "application/json")

	// Create the resonse
	resp := map[string]interface{}{
		"status":  true,
		"user":    *user,
		"token":   token,
		"message": "Successfully Created User"}

	json.NewEncoder(w).Encode(resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	var input UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	user.Name = input.Name

	models.DB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	models.DB.Delete(&user)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(user)
}

func GenerateBearerToken(user models.User) (string, error) {
	//Define the JWT with user information
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	}

	//Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Sign the token with Secret Key
	tokenString, err := token.SignedString([]byte("safety"))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
