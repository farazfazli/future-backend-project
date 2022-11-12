package futuredb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq" // required
)

var DBQueries *Queries
var db *sql.DB

func NewQueries() (*Queries, error) {
	if os.Getenv("POSTGRES_HOST") == "" ||
		os.Getenv("POSTGRES_PORT") == "" ||
		os.Getenv("POSTGRES_USERNAME") == "" ||
		os.Getenv("POSTGRES_PASSWORD") == "" || 
		os.Getenv("POSTGRES_DATABASE") == "" {
		return nil, errors.New("missing necessary variables")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USERNAME"),
	os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"))
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	DBQueries = New(db)
	return DBQueries, nil
}

func CloseDB() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}