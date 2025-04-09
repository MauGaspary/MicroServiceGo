package main

import (
	"fmt"
	"context"
	"github.com/MauGaspary/MicroServiceGo/application"
)


func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Falha ao iniciar o app:", err)
	}
}