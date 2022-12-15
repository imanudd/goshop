package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "goshop"
)

var psqlinfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var db *sql.DB
var err error

func Init() *sql.DB {
	db, err = sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatalf("Connection Fail !")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
