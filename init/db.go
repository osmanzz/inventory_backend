package init

import (
	"database/sql"
	"fmt"
	"log"
)

func NewDB() *sql.DB{
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "user=postgres password=osman dbname=osman sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Success!")
	}
	return db
}
