package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func main() {

	db, connectErr := connect("root", "password", "gatorrater")
	if connectErr != nil {
		log.Fatal(connectErr)
	}
	fmt.Printf("Connection Success!\n")

	var option = os.Args[1]

	var user = os.Args[2]
	var pass = os.Args[3]

	acc := Account{
		username: user,
		password: pass,
	}

	if option == "LOGIN" {
		fmt.Printf("Attempting login with user %v\n", acc.username)
		loginErr := login(db, acc)
		if loginErr != nil {
			log.Fatal(loginErr)
		}
		fmt.Printf("Succesfully logged in with user %v\n", acc.username)
	} else if option == "SIGNUP" {
		fmt.Printf("Attempting to create user %v\n", acc.username)
		signupErr := signup(db, acc)
		if signupErr != nil {
			log.Fatal(signupErr)
		}
		fmt.Printf("Succesfully created user %v\n", acc.username)
	} else if option == "DELETE" {
		fmt.Printf("Attempting to delete user %v\n", acc.username)
		deleteErr := delete(db, acc)
		if deleteErr != nil {
			log.Fatal(deleteErr)
		}
		fmt.Printf("Succesfully deleted user %v\n", acc.username)
	}
}
