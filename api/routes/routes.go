package routes

import (
	"main/api/controllers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/cars", controllers.HandleCars)
	mux.HandleFunc("/cars/", controllers.HandleCar)
	return mux
}
