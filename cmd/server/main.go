package main

import "fmt"

// Run instantiates and starts the application
func Run() error {
	fmt.Println("Starting the application...")
	return nil
}

func main() {
	fmt.Println("Go Rest API course")
	err := Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
