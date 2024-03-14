package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	user := flag.String("user", "root", "database user")
	password := flag.String("password", "root", "database password")
	dbname := flag.String("dbname", "test", "database name")
	host := flag.String("host", "localhost", "database host")
	crud := flag.String("crud", "create", "create, read, update, delete")

	flag.Parse()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", *user, *password, *host, *dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	switch *crud {
	case "create":
		//Create a new user
		_, err = db.Exec("INSERT INTO user(name,id) VALUES (?, ?)", "Jane Smith", 001)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("User was created successfully!")

	case "read":
		// Query all users from the "users" table
		rows, err := db.Query("SELECT id, name FROM user")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d: %s\n", id, name)
		}

	case "update":
		_, err = db.Exec("UPDATE user SET name = ? WHERE id = ?", "Jane Doe", 001)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("User was updated successfully!")
	case "delete":
		_, err = db.Exec("DELETE FROM user WHERE id = ?", 001)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("User was deleted successfully!")
	default:
		fmt.Println("Invalid CRUD operation")
	}
}
