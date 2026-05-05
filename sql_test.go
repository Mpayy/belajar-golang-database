package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	exec := "INSERT INTO customer (id, name) VALUES ('mita', 'Mita Selvira')"
	_, err := db.ExecContext(ctx, exec)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Customer")
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
