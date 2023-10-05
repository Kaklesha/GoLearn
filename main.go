package main

import (
   
    "fmt"
    "GoLearn/application"
    "context"
    
)

func main() {

    app:=application.New()
    err:= app.Start(context.TODO())
    if err != nil {
        fmt.Println("failed to start app: ",app) }
}