package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Age     int     `json:"age"`
	Balance int     `json:"balance"`
	Orders  []Order `json:"order"`
}

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UpdateStatus struct {
	Status string `json:"status"`
}

var users = []User{
	{ID: 1, Name: "Алескер", Email: "alesker@example.ru", Age: 24, Balance: 100500, Orders: []Order{}},
	{ID: 2, Name: "Аскеров", Email: "askerov@example.ru", Age: 25, Balance: 200600, Orders: []Order{}},
}

var orders = []Order{
	{ID: 1, UserID: 1, Name: "PlayStation 6", Status: "на складе"},
	{ID: 2, UserID: 2, Name: "Фигурка Человек-Паук в полный рост", Status: "в ПВЗ"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	user.ID = len(users) + 1
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	order.ID = len(orders) + 1
	orders = append(orders, order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func getUserOrders(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "некорректный ID"})
		return
	}

	for _, u := range users {
		if u.ID == userID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(u.Orders)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "пользователь не найден!"})
}

func updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "некорректный ID"})
		return
	}

	var updateStatus UpdateStatus
	if err := json.NewDecoder(r.Body).Decode(&updateStatus); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	for i, o := range orders {
		if o.ID == orderID {
			orders[i].Status = updateStatus.Status
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(orders[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "заказ не найден"})
}

func main() {

}
