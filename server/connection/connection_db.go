package connection

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDb() {

	var err error

	dsn := "root:5431sabi@tcp(127.0.0.1:3306)/grpc_crud_user"

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	log.Println("Database Succesfully Connected")
}
