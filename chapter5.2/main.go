package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:riiiight@/test?charset=utf8")
	checkErr(err)

	//insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("remenska", "icas", "2016-12-11")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//update
	stmt, err = db.Prepare("UPDATE userinfo SET username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("remenskaupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//query

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("uid=%d, username=%s, department=%s, created=%s \n", uid, username, department, created)
	}

	// delete
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("Rows affected:", affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
