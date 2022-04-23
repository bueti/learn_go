package main

import (
	"context"
	"fmt"

	"github.com/bueti/learn_go/rest-api/internal/db"
)

// is responsible for the instantiation
// and startup of our application
func Run() error {
	fmt.Println("starting up our application")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.Ping(context.Background()); err == nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
