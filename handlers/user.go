package handlers

import (
	//"fmt"
	"encoding/json"
	"net/http"

	"github.com/ibz11/go-restapi.git/models"

	//"github.com/ibz11/go-restapi.git/main"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
)

// type Repository struct {
// 	DB *gorm.DB
// }

// var p *Repository
// var DB *gorm.DB

func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.Users
	result := db.Find(&users)
	if result == nil {
		http.Error(w, "Users not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(users)

}
func GetAUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.Users
	// if params["id"] == nil{
	// 	http.Error(w, "No Id provided", http.StatusBadRequest)
	// }



	result := db.First(&user, params["id"])
	if result.Error != nil {
		// Check if the error is due to a record not found
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
   var params = mux.Vars(r)
    var user models.Users
	db.Delete(&user,params["id"])
	result := db.First(&user, params["id"])
	if result.Error != nil {
		// Check if the error is due to a record not found
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode("The User has been deleted successfully.")
}
func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
    var params = mux.Vars(r)
    var user models.Users
	//db.Update()
	db.Delete(&user,params["id"])
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON request", http.StatusBadRequest)
		return
	}
	result := db.Create(&user)
	if result.Error != nil {

		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)

}
