package main

import (
	"context"
	"fmt"

	"github.com/bueti/learn_go/rest-api/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
			Slug:   "manual-test",
			Author: "Ben",
			Body:   "Hello World",
		},
	)

	fmt.Println(cmtService.GetComment(context.Background(), "71c5d074-b6cf-11ec-b909-0242ac120002"))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
