package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	exec := "INSERT INTO user (username, password) VALUES ('admin', 'admin')"
	_, err := db.ExecContext(ctx, exec)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New User")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestMulti(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, create_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birthDate sql.NullTime
		var married bool
		var createAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("BirthDate:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Create_at:", createAt)
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "salah"

	login := "SELECT username FROM user WHERE username= ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, login, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login")
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "operator"
	password := "operator"

	exec := "INSERT INTO user (username, password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, exec, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New User")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "rifaih712@gmail.com"
	comment := "Test Komentar 3"

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert new comment with id", insertId)
}

func TestPrepareStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	prepare, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer prepare.Close()

	for i := 0; i < 10; i++ {
		email := "rifaih712" + strconv.Itoa(i) + "@gmail.com"
		comment := "Test Komentar" + strconv.Itoa(i)
		result, err := prepare.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin() // ini agar saat insert ke db tidak langsung auto commit
	if err != nil {
		panic(err)
	}
	defer tx.Rollback() // untuk mengirim data dan memastikan data yang masuk benar, tapi tidak di commit langsung
	// atau menghendle kalau ada kesalahan di bawah ini, maka tidak akan di commit, akan di rolback
	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "rifaih712" + strconv.Itoa(i) + "@gmail.com"
		comment := "Test Komentar" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", insertId)
	}
	err = tx.Commit() // untuk mengcommit data yang dikirm
	if err != nil {
		panic(err)
	}
}
