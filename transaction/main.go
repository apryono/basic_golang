package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	id    int
	name  string
	grade int
}

var dum = data{}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/book_store")
	handleError(err)

	defer db.Close()

	// tx artinya transaksi , variabel nya dapat diubah
	tx, err := db.Begin()
	handleError(err)

	// insert a record into table1
	res, err := tx.Exec("insert into customers(first_name, last_name, email) values (?,?,?)", "budi", "lawu", "budi@gmail.com")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// fetch the auto incremented id
	id, err := res.LastInsertId()
	fmt.Print(id)
	handleError(err)

	// insert record into table2, referencing the first record from table1
	res, err = tx.Exec("insert into orders(order_date,amount,customer_id) values (?,?,?)", "1998-02-19", 99.9, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// commit the transaction
	// handleError(tx.Commit())
	err = tx.Commit()
	if err != nil {
		fmt.Println("Something Problem, Please Check Your History")
	}

	log.Println("Done.")
}

func handleError(err error) {
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err.Error())
	}
}
