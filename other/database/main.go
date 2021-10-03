package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID int64
	Name string
	Age int64
}

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
	}

	//create table
	const sql = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	if _, err := db.Exec(sql); err != nil {
		fmt.Println(err)
	}

	//create record
	users := []*User{{Name: "tenntenn", Age: 42}, {Name: "Gopher", Age: 24}}
	for i := range users {
		const sql = "INSERT INTO user(name, age) values (?,?)"
		r, err := db.Exec(sql, users[i].Name, users[i].Age)
		if err != nil {fmt.Println(err)}
		id, err := r.LastInsertId()
		if err != nil {fmt.Println(err)}
		users[i].ID = id
		fmt.Println("INSERT", users[i])
	}

	//scan records
	age := 42
	rows, err := db.Query("SELECT * FROM user WHERE age = ?", age)
	if err != nil {fmt.Println(err)}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
	if err := rows.Err(); err != nil {fmt.Println(err)}

	//update record
	d, err := db.Exec("UPDATE user SET age = age + 1 WHERE id = 1")
	if err != nil {fmt.Println(err)}
	cnt, err := d.RowsAffected()
	if err != nil {fmt.Println(err)}
	fmt.Println("Affected rows:", cnt)
}
