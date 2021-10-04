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
	Amount int64
}

func CreateTable(db *sql.DB) {
	const sql = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		amount INTEGER NOT NULL
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
		Amount: int64(number),
	}}
	for i := range users {
		const sql = "INSERT INTO user(name, amount) values (?,?)"
		r, err := db.Exec(sql, users[i].Name, users[i].Amount)
		if err != nil {
			fmt.Println(err)
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}
		users[i].ID = id
		ScanRecords(db)
		fmt.Println("名前と金額をスペース区切りで入力")
	}
}

func ScanRecords(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Amount); err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("err: %v\n", err)
}

func Tranzaction(db *sql.DB, data []string) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	}

	name1, name2 := data[0], data[1]
	amount, err := strconv.Atoi(data[2])
	if err != nil {
		fmt.Println(err)
	}

	row1 := tx.QueryRow("SELECT * FROM user WHERE name = ?", name1)
	row2 := tx.QueryRow("SELECT * FROM user WHERE name = ?", name2)
	var u User
	if err := row1.Scan(&u.ID, &u.Name, &u.Amount); err != nil {
		tx.Rollback()
		fmt.Printf("err2: %v\n", err)
	}
	if err := row2.Scan(&u.ID, &u.Name, &u.Amount); err != nil {
		tx.Rollback()
		fmt.Printf("err2: %v\n", err)
	}

	const updateSQL = "UPDATE user SET amount = ? WHERE name = ?"
	if _, err = tx.Exec(updateSQL, u.Amount+int64(amount), name2); err != nil {
		tx.Rollback()
		fmt.Printf("err3: %v\n", err)
	}
	if err := tx.Commit(); err != nil {
		fmt.Printf("err4: %v\n", err)
	}
}

func main() {
	/* open database */
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
	}

	ScanRecords(db)
	// CreateTable(db)
	// var record [2]string = [2]string{"huga", "0"}
	// CreateRecord(db, record[:])

	fmt.Println("名前と番号をスペース区切りで入力")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Tranzaction(db, strings.Split(scanner.Text(), " "))
		// CreateRecord(db, strings.Split(scanner.Text(), " "))
	}
}
