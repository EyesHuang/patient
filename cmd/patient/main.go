package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	rDB "manage-patinets/db"
	rHttp "manage-patinets/http"

	"github.com/joho/godotenv"
)

const (
	HOST = "localhost"
	PORT = "5432"
	SSL  = "disable"
)

func main() {
	port := flag.Int("port", 8080, "listen port")
	flag.Parse()
	if err := run(*port); err != nil {
		fmt.Println("error to run service")
	}
}

func run(port int) error {
	repo := rDB.NewPatientRepository(openDB())
	srv := rHttp.NewServer(repo)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), srv)
}

func openDB() *sql.DB {
	err := godotenv.Load("./deployment/postgresql/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the environment variables
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPasswd := os.Getenv("POSTGRES_PASSWORD")

	driver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, dbUser, dbPasswd, dbName, SSL)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic("open database error")
	}
	return db
}
