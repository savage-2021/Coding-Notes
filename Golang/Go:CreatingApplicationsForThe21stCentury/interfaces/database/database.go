package database

import "fmt"

type Database struct {
	pool string
}

func Start() *Database{
	database := &Database{}
	database.pool = "examplepool"
	return database
}

func(db *Database) Select() {
	fmt.Println("Running a Select")
}

func(db *Database) Insert() {
	fmt.Println("Running an Insert")
}

func (db *Database) Delete() {
	fmt.Println("Running a Delete")
}