package MPPL

import (
"MPPL/controllers"
"MPPL/middlewares"
"MPPL/repository"
"github.com/gorilla/mux"
"github.com/jinzhu/gorm"

)

func LoadRouter(db *gorm.DB) (r *mux.Router) {
	userRepo := repository.newBrandsRepo(db)
	userController := controllers.NewUserController(userRepo)

	r = mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/users", userController.Resources).Methods("GET", "POST")
	v1.HandleFunc("/users/{id}", userController.Resources).Methods("GET", "PATCH", "DELETE")

	r.Use(middlewares.LoggerMidldlware)

	return
}
