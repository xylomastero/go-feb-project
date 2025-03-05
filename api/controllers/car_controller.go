package controllers

import (
	"encoding/json"
	"main/api/db"
	"main/api/models"
	"net/http"
	"strconv"
	"strings"
)

func HandleCars(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCars(w, r)
	case "POST":
		AddCar(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleCar(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
	switch r.Method {
	case "GET":
		GetCar(w, r, id)
	case "PUT":
		UpdateCar(w, r, id)
	case "DELETE":
		DeleteCar(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func GetCars(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, make, model, year FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	json.NewEncoder(w).Encode(cars)
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO cars (make, model, year) VALUES (?, ?, ?)", car.Make, car.Model, car.Year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	car.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func GetCar(w http.ResponseWriter, r *http.Request, id string) {
	carID, _ := strconv.Atoi(id)
	var car models.Car
	err := db.DB.QueryRow("SELECT id, make, model, year FROM cars WHERE id = ?", carID).Scan(&car.ID, &car.Make, &car.Model, &car.Year)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request, id string) {
	carID, _ := strconv.Atoi(id)
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE cars SET make = ?, model = ?, year = ? WHERE id = ?", car.Make, car.Model, car.Year, carID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	car.ID = carID
	json.NewEncoder(w).Encode(car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request, id string) {
	carID, _ := strconv.Atoi(id)
	_, err := db.DB.Exec("DELETE FROM cars WHERE id = ?", carID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
