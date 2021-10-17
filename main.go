package main

import (
	"yofio/avanzado/app"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app.StartApp()
}
