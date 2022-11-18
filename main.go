package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/vldcreation/simple-auth-golang/internal/app"
)

func main() {
	godotenv.Load()

	app.Run(context.Background())
}
