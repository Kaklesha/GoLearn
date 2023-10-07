package main

import (
	"GoLearn/application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//defer database.Close()
	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
