package main

import "fmt"

// is responsible for the instantiation
// and startup of our application
func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
