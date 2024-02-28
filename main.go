package main

import (
	"fmt"
	"log"
	"net/http"
	

	"github.com/gorilla/mux"
	//"gorm.io/gorm"
	"github.com/ibz11/go-restapi.git/config"
	"github.com/ibz11/go-restapi.git/handlers"
	
)



func initializeRouter() {
	db,err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	app := mux.NewRouter()

	//Routes
	//r.PathPrefix("/users/")
	route := app.PathPrefix("/api").Subrouter()
	route.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsers(w, r, db)
	}).Methods("GET")

	route.HandleFunc("/users/{id}",func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAUser(w, r, db)
	}).Methods("GET")


	route.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, db)
	}).Methods("POST")




	route.HandleFunc("/users/{id}",func(w http.ResponseWriter, r *http.Request) {
	          handlers.UpdateUser(w, r, db)
	}).Methods("PUT")


	route.HandleFunc("/users/{id}",func(w http.ResponseWriter, r *http.Request) {
	            handlers.DeleteUser(w,r,db)
	}).Methods("DELETE")


	fmt.Println("Starting server at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", route))

}



func main() {
	
	initializeRouter()


}
