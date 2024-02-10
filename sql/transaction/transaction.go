package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("Insert into usuarios(id, nome) values(?,?)")
	stmt.Exec(200, "Maria")
	stmt.Exec(199, "João")
	_ , err = stmt.Exec(199, "Pedro") 
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}