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

func GetJwtSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}
func GetAdminID() string {
	return os.Getenv("ADMIN_ID")
}
func GetManagerID() string {
	return os.Getenv("MANAGER_ID")
}
func GetUserID() string {
	return os.Getenv("USER_ID")
}
