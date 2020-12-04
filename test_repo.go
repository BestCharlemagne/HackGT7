package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type TestRepo struct {
	database *sql.DB
}

func (repo TestRepo) GetDatabase() *sql.DB {
	if repo.database == nil {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/store_navigator")

		if err != nil {
			log.Fatal(err)
		}
		repo.database = db
	}
	return repo.database
}
