package ui

import (
	"fmt"
)

func ShowMenu() {
	fmt.Println("1 - Create account")
	fmt.Println("2 - List accounts")
	fmt.Println("3 - Transfer")
	fmt.Println("4 - Withdraw")
	fmt.Println("5 - Deposit")
	fmt.Println("6 - Remove account")
	fmt.Println("7 - Exit")
}

func ShowCreateAccountMenu() {
	fmt.Println("Creating account")
}