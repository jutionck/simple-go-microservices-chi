package main

import (
	"context"
	"fmt"

	"mipdevp.com/golang-microservices/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start serve", err)
	}
}
