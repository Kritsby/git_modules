package db

import "fmt"

type DB struct{}

func (db *DB) ConnectDB() {
	fmt.Println("u are connected now!")
}
