package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

// func linearSearch(arr []string, s string) int {
//     for i, v := range arr {
//         if s == v {
//             return i
//         }
//     }

//     return -1
// }

func SearchForName(username string, users Users) int {
	for i, v := range users.Users {
		if username == v.Username {
			return i
		}
	}
	return -1
}

func SearchForPassword(password string, users Users) int {
	for i, v := range users.Users {
		if password == v.Password {
			return i
		}
	}
	return -1
}

func main() {
	file, _ := ioutil.ReadFile("data.json")

	data := Users{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Users); i++ {
		fmt.Println("Name: ", data.Users[i].Name)
		fmt.Println("Username: ", data.Users[i].Username)
		fmt.Println("Type: ", data.Users[i].Type)
		fmt.Println("Age: ", data.Users[i].Age)
		fmt.Println("Password: ", data.Users[i].Password)
	}

	fmt.Println("")
	var inputUserName string
	fmt.Print("Enter a username: ")
	fmt.Scanln(&inputUserName)

	fmt.Println("")
	fmt.Print("Enter a password: ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	if Auth(inputUserName, inputPassword, data) {
		fmt.Println("You are logged in!")
	} else {
		fmt.Println("You are not logged in!")
	}

}

func Auth(username string, password string, users Users) bool {
	if SearchForName(username, users) != -1 && SearchForPassword(password, users) != -1 {
		return true
	}
	return false
}
