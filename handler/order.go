package handler

import (
	"GoLearn/model"
	"GoLearn/repository/order"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Order struct {
	Repo *order.SqlRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Create an order")
	var body struct {
		CustomerID uint             `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	now := time.Now().UTC()
	order := model.Order{
		CustomerID: body.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}
	err := o.Repo.Insert(order)
	if err != nil {
		fmt.Println("failed to isert:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
	//2.48
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List an order")

}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order ID")

}
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order ID")

}
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order ID")

}
