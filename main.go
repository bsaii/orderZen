package main

import (
	"context"
	"fmt"

	"github.com/bsaii/orderZen/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
