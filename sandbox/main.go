package main

import (
	"fmt"
	"log"

	database "github.com/pyumz/social/internal"
)

func main() {

	c := database.NewClient("db.json")
	err := c.EnsureDB()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database was created and/or already exists!")

}
