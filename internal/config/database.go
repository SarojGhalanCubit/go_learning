
package config

import (
	"context"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(),
		os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	
	// ✅ Print success message
	log.Println("✅ Database connected successfully")
	return conn
}
