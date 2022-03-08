package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

type User struct {
	Id         int
	Name       string
	Username   string
	Email      string
	Password   string
	Created_at string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "users",
	}
	fmt.Println(cfg.FormatDSN())
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "root:terminator21@/users")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	users, err := UserByName("Mimmo Baluyut")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User found: %v\n", users)
}

func UserByName(name string) ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT * FROM user WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Username, &usr.Email, &usr.Password, &usr.Created_at); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return users, nil
}
