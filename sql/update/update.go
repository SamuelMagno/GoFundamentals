package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	st.Exec("Samuel", 1)
	st.Exec("Júlia", 2)
}