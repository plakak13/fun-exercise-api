package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func New() (*Postgres, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("HOST")
	user := os.Getenv(("USERDB"))
	pass := os.Getenv("PASS")
	dbName := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}

	databaseSource := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, portInt, user, pass, dbName)
	db, err := sql.Open("postgres", databaseSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &Postgres{Db: db}, nil
}
