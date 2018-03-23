package main

import (
	"github.com/maikeulb/national-parks/app"
	"os"
)

func main() {
	a := app.App{}
	a.Initialize(
		// "172.17.0.2",
		// 5432,
		// "postgres",
		// "P@ssw0rd!",
		// "national_parks")
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DB"))

	a.Run(":5000")
}
