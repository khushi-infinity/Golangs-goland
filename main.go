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

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance $%d) : ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be greater than your balance")
		} else {
			break
		}
	}
	return bet
}

func main() {
	symbols := map[string]uint{
		"A": 3,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	balance := uint(200)
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet

	}

	fmt.Printf("You have left with, $%d.\n", balance)
}
