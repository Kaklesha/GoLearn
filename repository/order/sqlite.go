package order

import (
	"GoLearn/model"
	"database/sql"
	"fmt"
)

type SqlRepo struct {
	DB *sql.DB
}

func orderIDKey(id uint) string {
	return fmt.Sprintf("order:%d", id)

}

func (r *SqlRepo) Insert(order model.Order) error {
	statement, err := r.DB.Prepare(`INSERT INTO order (
			customer_id,
			line_items,
			created_at,
			shipped_at,
			completed at) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert order: %w", err)
	}
	_, err = statement.Exec(order.CustomerID, order.LineItems, order.CreatedAt, order.CreatedAt, order.ShippedAt, order.CompletedAt)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}
	return nil
}

func (r *SqlRepo) FindById(order_id uint) (model.Order, error) {
	model_order := model.Order{}
	line_item := ""
	rows, err := r.DB.Query(`SELECT * FROM order WHERE order_id =(?))`, order_id)
	if err != nil {
		return model_order, fmt.Errorf("failed to prepare insert order: %w", err)
	}
	rows.Scan(&model_order.OrderID, &model_order.CustomerID, &line_item, &model_order.CreatedAt, &model_order.ShippedAt, &model_order.CompletedAt)

	model_order.UnmarshalLineItems(line_item)
	return model_order, nil
}

func (r *SqlRepo) DeleteByID(order_id uint) error {

	_, err := r.DB.Query(`DELETE FROM order WHERE order_id =(?))`, order_id)
	if err != nil {
		return fmt.Errorf("failed to prepare delete order: %w", err)
	}
	return nil
}

// func (r *SqlRepo) UpdateById(colunm model.Order, value string, order_id uint) (model.Order, error) {
// 	model_order := model.Order{}
// //	line_item := ""
// 	rows, err := r.db.Query(`UPDATE ? FROM order WHERE order_id =(?))`)
// 	if err != nil {
// 		return model_order, fmt.Errorf("failed to prepare insert order: %w", err)
// 	}
// 	rows.Scan(,&model_order.OrderID)

// 	model_order.UnmarshalLineItems(line_item)
// 	return model_order, nil
// }

// func (r *SqlRepo) UpdateById(colunm model.Order, value string,order_id uint) error {

// 	_, err := r.db.Query(`UPDATE order SET ?=? WHERE order_id =(?))`,colunm, value, order_id)
// 	if err != nil {
// 		return fmt.Errorf("failed to prepare update order: %w", err)
// 	}
// 	return nil
// }

// UPDATE имя_таблицы
// SET столбец1 = значение1, столбец2 = значение2, ... столбецN = значениеN
// [WHERE условие_обновления]

func (r *SqlRepo) Update(model_order model.Order) error {

	statement, err := r.DB.Prepare(`UPDATE order SET
        customer_id = ?,
        line_items = ?,
        created_at = ?,
        shipped_at = ?,
        completed at = ?
        WHERE order_id = ?`)
	if err != nil {
		return fmt.Errorf("failed to prepare Update order: %w", err)
	}

	_, err = statement.Exec(model_order.CustomerID, model_order.MarshalLineItems(), model_order.CreatedAt, model_order.CreatedAt, model_order.ShippedAt, model_order.CompletedAt, model_order.OrderID)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}
	return nil

}

func (r *SqlRepo) FindAll() ([]model.Order, error) {
	array := []model.Order{}
	model_order := model.Order{}
	line_item := ""
	rows, err := r.DB.Query(`SELECT * FROM order`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare insert order: %w", err)
	}
	for rows.Next() {
		rows.Scan(&model_order.OrderID, &model_order.CustomerID, &line_item, &model_order.CreatedAt, &model_order.ShippedAt, &model_order.CompletedAt)
		model_order.UnmarshalLineItems(line_item)
		array = append(array, model_order)
	}

	return array, nil
}

// func main() {
//     db, err := sql.Open("sqlite3", "./database.db")
//     if err != nil {
//         panic(err)
//     }

//     defer db.Close()

//     statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS time (id INTEGER PRIMARY KEY, time DATETIME)")
//     if err != nil {
//         panic(err)
//     }
//     statement.Exec()

//     statement, err = db.Prepare("INSERT INTO time (time) VALUES (?, ?)")
//     if err != nil {
//         panic(err)
//     }
//     statement.Exec(time.Now().Add(time.Hour * 2), time.Now())

//     rows, _ := db.Query("SELECT id, time FROM time")
//     var id int
//     var cTime time.Time

//     for rows.Next() {
//         rows.Scan(&id, &cTime)
//         fmt.Println(id, cTime)
//     }
// }
