package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Login(1) or create user(2)?")
	fmt.Print("| => ")
	// var choice int
	scanner.Scan()
	choice := scanner.Text()
	// fmt.Scanln(&choice)
	if choice == "1" {
		fmt.Println("Enter Username:")
		fmt.Print("| => ")
		// var username string
		// fmt.Scanln(&username)
		scanner.Scan()
		username := scanner.Text()

		fmt.Println("Enter Password:")
		fmt.Print("| => ")
		// var password string
		// fmt.Scanln(&password)
		scanner.Scan()
		password := scanner.Text()

		if login(username, password) != nil {
			fmt.Println("Login success")
			maininterface()
		} else {
			fmt.Println("Login failed")
		}
	} else if choice == "2" {
		fmt.Println("Enter Name:")
		fmt.Print("| => ")
		// var Name string
		// fmt.Scanln(&Name)
		scanner.Scan()
		Name := scanner.Text()

		fmt.Println("Enter Username:")
		fmt.Print("| => ")
		// var Username string
		// fmt.Scanln(&Username)
		scanner.Scan()
		Username := scanner.Text()

		fmt.Println("Enter Email:")
		fmt.Print("| => ")
		// var Email string
		// fmt.Scanln(&Email)
		scanner.Scan()
		Email := scanner.Text()

		fmt.Println("Enter Password:")
		fmt.Print("| => ")
		// var Password string
		// fmt.Scanln(&Password)
		scanner.Scan()
		Password := scanner.Text()

		usrID, err := createUser(User{
			Name:     Name,
			Username: Username,
			Email:    Email,
			Password: Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID of added User: %v\n", usrID)
		if login(Username, Password) != nil {
			fmt.Println("Sign Up and Login success")
			maininterface()
		} else {
			fmt.Println("Sign Up and Login failed")
		}
	}
}

func createUser(usr User) (int64, error) {
	result, err := db.Exec("INSERT INTO user (name, username, email, password) VALUES (?, ?, ?, ?)", usr.Name, usr.Username, usr.Email, usr.Password)
	if err != nil {
		return 0, fmt.Errorf("createUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createUser: %v", err)
	}
	return id, nil
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

func maininterface() {

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
