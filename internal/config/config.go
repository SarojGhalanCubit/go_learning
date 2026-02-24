package config

import "os"

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8080"
	}

	return ":" + port
}
