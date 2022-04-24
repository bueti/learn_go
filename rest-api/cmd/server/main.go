package main

import (
	"fmt"

	"github.com/bueti/learn_go/rest-api/internal/comment"
	"github.com/bueti/learn_go/rest-api/internal/db"
	transportHttp "github.com/bueti/learn_go/rest-api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
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
