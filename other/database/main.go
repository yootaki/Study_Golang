package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID     int64
	Name   string
	Number int64
}

func CreateTable(db *sql.DB) {
	const sql = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		number INTEGER NOT NULL
	);
	`
	if _, err := db.Exec(sql); err != nil {
		fmt.Println(err)
	}
}

func CreateRecord(db *sql.DB, data []string) {
	number, err := strconv.Atoi(data[1])
	if err != nil {
		fmt.Println(err)
	}
	users := []*User{{
		Name:   data[0],
		Number: int64(number),
	}}
	for i := range users {
		const sql = "INSERT INTO user(name, number) values (?,?)"
		r, err := db.Exec(sql, users[i].Name, users[i].Number)
		if err != nil {
			fmt.Println(err)
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}
		users[i].ID = id
		ScanRecords(db)
		fmt.Println("名前と番号をスペース区切りで入力")
	}
}

func ScanRecords(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Number); err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("err: %v\n", err)
}

func main() {
	/* open database */
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
	}

	ScanRecords(db)
	// CreateTable(db)

	fmt.Println("名前と番号をスペース区切りで入力")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		CreateRecord(db, strings.Split(scanner.Text(), " "))
	}
}
