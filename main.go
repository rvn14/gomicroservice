package main

import (
    "context"
    "fmt"
    "github.com/rvn14/gomicroservice/application"
)

func main() {
    app:= application.New()

    err := app.Start(context.TODO())
    if err != nil {
        fmt.Printf("Failed to start application: %v\n", err)
    }

}