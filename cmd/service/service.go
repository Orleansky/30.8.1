package main

import (
	"Anastasia/skillfactory/databases/module30/pkg/storage"
	"Anastasia/skillfactory/databases/module30/pkg/storage/postgres"
	"fmt"
	"log"
)

var db storage.Interface

func main() {
	var err error

	connstr :=
		"postgres://postgres:qwerty@localhost/tasks?sslmode=disable"

	db, err = postgres.New(connstr)

	if err != nil {
		log.Fatal(err)
	}

	tasks, err := db.GetTasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
}
