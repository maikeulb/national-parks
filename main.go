package main

func main() {
	a := App{}
	a.Initialize("172.17.0.2", 5432, "postgres", "P@ssw0rd!", "national_parks")
	// os.Getenv("APP_DB_HOST"),
	// os.Getenv("APP_DB_PORT"),
	// os.Getenv("APP_DB_USERNAME"),
	// os.Getenv("APP_DB_PASSWORD"),

	a.Run(":8080")
}
