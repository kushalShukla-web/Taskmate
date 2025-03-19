package product

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "kushal"
	password = "mysecretpassword"
	dbname   = "codexray"
)

func InitDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err := Db.Ping(); err != nil {
		return err
	}

	log.Println("Successfully Connected to Database")
	return nil
}
func GetDb() *sql.DB {
	return Db
}
