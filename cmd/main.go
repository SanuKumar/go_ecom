package main

import (
	"database/sql"
	"github/SanuKumar/go_ecom/cmd/api"
	"github/SanuKumar/go_ecom/config"
	"github/SanuKumar/go_ecom/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	initStorage(db)
	server := api.NewAPIServer(":8080", nil) // Replace nil with actual database connection
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error pinging MySQL: %v", err)
	}
	log.Println("Successfully connected to MySQL database")

}
