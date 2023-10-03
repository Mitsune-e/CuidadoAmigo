package DbConnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("", "")

	errCheck(err)
	err = db.Ping()
	errCheck(err)
	defer db.Close()
	fmt.Println("Connection worked")
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
