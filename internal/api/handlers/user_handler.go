package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Pramessh-P/weather-api/internal/database"
	"github.com/Pramessh-P/weather-api/internal/helper"
	"github.com/Pramessh-P/weather-api/internal/model"
	"github.com/Pramessh-P/weather-api/internal/service"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserhandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "can't access with this requst", http.StatusNotAcceptable)
		return
	}
	var user *model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "please fill form correctly", http.StatusBadRequest)
		return
	}

	err := u.userService.Signup(user)
	if err != nil {
		http.Error(w, "can't signup "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": "user signup successfully with " + user.Email,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user helper.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "User form is required", http.StatusInternalServerError)
		return
	}
	err := u.userService.Login(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": "user login successfully with " + user.Email,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func (u *UserHandler) GetRecentAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method restricted!!", http.StatusBadRequest)
		return
	}
	userEmail := r.FormValue("email")
	err := u.userService.GetRecentActivities(userEmail)
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "user is not logined.. please login.", http.StatusBadRequest)
		return
	}
	var results []model.UserActivity
	err = database.DB.Where("email = ?",userEmail ).Order("id desc").Limit(10).Find(&results).Error
	if err != nil {
		http.Error(w,"error from retrieve data",http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write JSON data as the response
		_, err = w.Write(jsonData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}
