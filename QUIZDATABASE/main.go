package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type question struct {
	Id          int
	question    string
	answer      string
	explination string
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "users",
	}
	fmt.Println(cfg.FormatDSN())

	var err error
	db, err = sql.Open("mysql", "root:terminator21@/quiz")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	questions := getQuestions()
	scanner := bufio.NewScanner(os.Stdin)
	correct := 0
	for _, val := range questions {
		fmt.Printf("%v\n", val.question)
		fmt.Print("| => ")
		scanner.Scan()
		answer := scanner.Text()
		if hash(answer) == val.answer {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Wrong!")
			fmt.Println(val.explination)
		}
	}
	fmt.Printf("You got %v out of %v questions correct!\n", correct, len(questions))
}

func getQuestions() []question {
	var questions []question

	rows, err := db.Query("SELECT * FROM 4aq")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var qst question
		if err := rows.Scan(&qst.Id, &qst.question, &qst.answer, &qst.explination); err != nil {
			return nil
		}
		questions = append(questions, qst)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return questions

}
