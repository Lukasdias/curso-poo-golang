package main

import (
	"fmt"
	"math/rand"
	"time"

	"poogolang.com/m/v2/accounts"
	"poogolang.com/m/v2/api"
	"poogolang.com/m/v2/clients"
	"poogolang.com/m/v2/ui"
)



func main () {
	checkingAccounts := *new([]accounts.CheckingAccount)
	savingsAccounts := *new([]accounts.SavingsAccount)
	option := -1

	fmt.Println("Buscando dados ficticios...")
	onInit(&checkingAccounts, &savingsAccounts)

	fmt.Println("Dados carregados com sucesso!")

	for option != 0 {
		ui.ShowMenu()
		option = ui.ReadOption()
		handleOptions(option, 	&checkingAccounts, &savingsAccounts)
	}
}

func onInit(checkingAccounts *[]accounts.CheckingAccount, savingsAccounts *[]accounts.SavingsAccount) {
	db := api.GetRandomPeople()
	for _, person := range db {
		holder := clients.Holder{
			Name: person.Name.First + " " + person.Name.Last,
			Cpf: "123.456.789-10",
			Profission: "Developer",
			Email: person.Email,
			Gender: person.Gender,
			PhoneNum: person.Phone,
		}
		checkingAccount := accounts.CheckingAccount{
			Account: accounts.Account{
				Holder: holder,
				NumAgency: 123,
				NumAccount: 123456,
				Funds: float64(generateRandomNumber(100, 1000000)),
				Email: person.Email,
			},
		}
		savingsAccount := accounts.SavingsAccount{
			Account: accounts.Account{
				Holder: holder,
				NumAgency: 123,
				NumAccount: 123456,
				Funds: float64(generateRandomNumber(100, 1000000)),
				Email: person.Email,
			},
		}
		*checkingAccounts = append(*checkingAccounts, checkingAccount)
		*savingsAccounts = append(*savingsAccounts, savingsAccount)
	}
}

func handleOptions(option int, checkingAccounts *[]accounts.CheckingAccount, savingsAccounts *[]accounts.SavingsAccount,
	) {
	switch option {
	case 1:
		ui.ShowListAccountsMenu(
			checkingAccounts,
			savingsAccounts,
		)
	case 0:
		fmt.Println("Saindo...")
	default:
		fmt.Println("Invalid option!")
	}
}

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
