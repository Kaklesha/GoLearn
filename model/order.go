package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Order struct {
	OrderID     uint       `json:"order_id"`
	CustomerID  uint       `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
type LineItem struct {
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
	Price    uint `json:"price"`
}

func (o *Order) MarshallLineItems() string {
	res := ""
	for index, item := range o.LineItems {
		if index != 0 {
			res += fmt.Sprintf(",%v", item.ItemID)
		} else {
			res += fmt.Sprint(item.ItemID)
		}
	}

	return res
}

func (o *Order) UnmarshallLineItems(ids string, db *sql.DB) {

	rows, _ := db.Query("SELECT * FROM line_item WHERE item_id IN (?)", items)
	item := LineItem{}

	for rows.Next() {
		rows.Scan(&item.ItemID, &item.Price, &item.Quantity)
		o.LineItems = append(o.LineItems, item)
	}
}
