package main

import "fmt"

func getName() string {
	name := ""
	fmt.Println("Welcome to the CASINO!")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, Let's Play!\n", name)
	return name
}

func main() {
	getName()
}
