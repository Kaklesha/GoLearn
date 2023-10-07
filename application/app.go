package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	router http.Handler
	db     *sql.DB
}

// func connectWithBD() *sql.DB {
// 	database, err := sql.Open("sqlite3", "./database.db")
//     if err != nil {
// 		log.Fatal(err)
// 	}
// 	return database
// }

func New() *App {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	app := &App{
		router: loadRoutes(),
		db:     db,
	}
	return app
}

// func New() *App{
// 	app:= & App{
// 		router: loadRoutes(),
// 		db: connectWithBD()
// 	}
// 	return app
// }

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to starting server: %w", err)
	}

	statement, err := a.db.Prepare(`CREATE TABLE IF NOT EXISTS line_item (
		item_id INTEGER PRIMARY KEY,
		 quantity INTEGER,
		 price INTEGER)`)

	if err != nil {
		panic(err)
	}
	statement.Exec()

	// statement, err = a.db.Prepare("INSERT INTO time (time) VALUES (?)")
	// if err != nil {
	//     panic(err)
	// }

	statement, err = a.db.Prepare(`CREATE TABLE IF NOT EXISTS order (
        order_id INTEGER PRIMARY KEY, 
        customer_id INTEGER, 
        line_items TEXT, 
        created_at DATETIME,
        shipped_at DATETIME,
        completed_at DATETIME)`)
	if err != nil {
		panic(err)
	}
	statement.Exec()

	// statement.Exec(time.Now().Add(time.Hour * 2))

	// rows, _ := db.Query("SELECT id, time FROM time")
	// var id int
	// var cTime time.Time

	// for rows.Next() {
	//     rows.Scan(&id, &cTime)
	//     fmt.Println(id, cTime)
	// }

	defer func() {
		if err := a.db.Close(); err != nil {
			fmt.Println("failed to close db", err)
		}
	}()

	fmt.Println("starting server")

	ch := make(chan error, 1) //ch канал для async chan - канал error - тип данных 1 - размер буфера

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to listen to server: %w", err)
		}
	}()
	select {

	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
	return nil
}
