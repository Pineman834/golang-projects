package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
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
	// users, err := UserByName("Mimmo Baluyut")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User found: %v\n", users)

	fmt.Println("Login(1) or create user(2)?")
	fmt.Print("| => ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Println("Enter Username:")
		fmt.Print("| => ")
		var username string
		fmt.Scanln(&username)

		fmt.Println("Enter Password:")
		fmt.Print("| => ")
		var password string
		fmt.Scanln(&password)

		if login(username, password) != nil {
			fmt.Println("Login success")
		} else {
			fmt.Println("Login failed")
		}
	} else if choice == 2 {
		// TODO: Create user
	}
}

func login(liveusername string, livepassword string) []User {
	var users []User

	rows, err := db.Query("SELECT * FROM user WHERE username = ? AND password = ?", liveusername, livepassword)
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Username, &usr.Email, &usr.Password, &usr.Created_at); err != nil {
			return nil
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return users
}

// func UserByName(name string) ([]User, error) {
// 	var users []User

// 	rows, err := db.Query("SELECT * FROM user WHERE name = ?", name)
// 	if err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var usr User
// 		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Username, &usr.Email, &usr.Password, &usr.Created_at); err != nil {
// 			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 		}
// 		users = append(users, usr)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	return users, nil
// }
