package repository

import (
	"database/sql"
	"encoding/json"
	"log"

	//mysql driver is used to setup the sql methods to work with a mysql server
	_ "github.com/go-sql-driver/mysql"
)

//ExampleRepo is a struct to represent a testing MySQL database
type ExampleRepo struct {
	database *sql.DB
}

//Scan overrides the sql.Scan method to parse the json from the database to a PathType
func (m *PathType) Scan(src interface{}) error {
	// The data stored in a JSON field is actually returned as []uint8
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}

//Scan overrides the sql.Scan method to parse the json from the database to a PathType
func (m *ItemType) Scan(src interface{}) error {
	// The data stored in a JSON field is actually returned as []uint8
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}

//GetAllStores currently retieves all stores with only support for a test SQL database.
func (repo ExampleRepo) GetAllStores() []Store {
	rows, err := repo.getDatabase().Query("SELECT * FROM stores")
	defer rows.Close()

	var stores []Store

	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			var newStore Store
			err := rows.Scan(&newStore.ID, &newStore.Title, &newStore.ZipCode, &newStore.Path, &newStore.Items)
			if err != nil {
				log.Fatal(err)
			}
			stores = append(stores, newStore)
		}
	}
	return stores
}

//AddStore adds a single store to the database
func (repo ExampleRepo) AddStore(store Store) {
	database := repo.getDatabase()
	pathJSON, pathErr := json.Marshal(store.Path)
	itemsJSON, itemsErr := json.Marshal(store.Items)
	if pathErr == nil && itemsErr == nil {
		stmt, prepareErr := database.Prepare("INSERT INTO stores VALUES (?, ? , ?, ?, ?)")
		defer stmt.Close()

		if prepareErr != nil {
			log.Fatal(prepareErr)
		}

		result, executeErr := stmt.Exec(store.ID, store.Title, store.ZipCode, pathJSON, itemsJSON)

		if executeErr == nil {
			print(result.RowsAffected())
		} else {
			log.Fatal(executeErr)
		}
	} else {
		if pathErr != nil {
			log.Fatal(pathErr)
		}
		if itemsErr != nil {
			log.Fatal(itemsErr)
		}
	}
}

//AddAllStores adds a list of stores to the database
func (repo ExampleRepo) AddAllStores(stores []Store) {
	database := repo.getDatabase()
	tx, txErr := database.Begin()
	if txErr == nil {
		stmt, prepareErr := database.Prepare("INSERT INTO stores VALUES (?, ? , ?, ?, ?)")
		defer stmt.Close()

		if prepareErr == nil {
			for _, store := range stores {
				pathJSON, pathErr := json.Marshal(store.Path)
				itemsJSON, itemsErr := json.Marshal(store.Items)
				if pathErr == nil && itemsErr == nil {

					result, executeErr := stmt.Exec(store.ID, store.Title, store.ZipCode, pathJSON, itemsJSON)

					if executeErr == nil {
						print(result.RowsAffected())
					} else {
						log.Fatal(executeErr)
					}
				} else {
					if pathErr != nil {
						log.Fatal(pathErr)
					}
					if itemsErr != nil {
						log.Fatal(itemsErr)
					}
				}
			}
			tx.Commit()
		} else {
			log.Fatal(prepareErr)
		}
	} else {
		log.Fatal(txErr)
	}
}

//RemoveStore removes a specific store, decided by its id
//Be careful, if you have multiple stores with same id, all will be deleted.
func (repo ExampleRepo) RemoveStore(store Store) {
	database := repo.getDatabase()
	stmt, err := database.Prepare("DELETE FROM stores WHERE id=?")
	defer stmt.Close()
	if err == nil {
		result, execError := stmt.Exec(store.ID)

		if execError != nil {
			log.Fatal(execError)
		} else {
			print(result.RowsAffected())
		}
	} else {
		log.Fatal(err)
	}
}

//RemoveAllStores removes all stores from the database
func (repo ExampleRepo) RemoveAllStores() {
	database := repo.getDatabase()
	stmt, err := database.Prepare("DELETE FROM stores")
	defer stmt.Close()
	if err == nil {
		result, execError := stmt.Exec()

		if execError != nil {
			log.Fatal(execError)
		} else {
			print(result.RowsAffected())
		}
	} else {
		log.Fatal(err)
	}
}

//Close closes the database connection
func (repo ExampleRepo) Close() {
	repo.database.Close()
}

func (repo ExampleRepo) getDatabase() *sql.DB {
	if repo.database == nil {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/store_navigator")

		if err != nil {
			log.Fatal(err)
		}
		repo.database = db
	}
	return repo.database
}
