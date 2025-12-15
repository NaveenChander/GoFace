package main

import (
	"io"
	"fmt"
	
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello, World!")
	// Example of making a simple HTTP GET request
	// to demonstrate functionality.
	makeGetRequest()
}

func makeGetRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("API Response:", string(body))
}

// createPostgresConnection establishes a connection to the PostgreSQL database.
func createPostgresConnection(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// insertRow inserts a row into the specified table with given columns and values.
func insertRow(db *sql.DB, table string, columns []string, values []interface{}) error {
	placeholders := ""
	for i := range columns {
		if i > 0 {
			placeholders += ", "
		}
		placeholders += fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table,
		joinColumns(columns), placeholders)
	_, err := db.Exec(query, values...)
	return err
}

// joinColumns joins column names with commas for SQL queries.
func joinColumns(columns []string) string {
	return fmt.Sprintf("%s",
		stringJoin(columns, ", "))
}

// stringJoin joins a slice of strings with a separator.
func stringJoin(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
