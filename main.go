package main

import (
	"fmt"
)

func main () {
	accounts := make(map[string]CheckingAccount)
	accounts["123456"] = CheckingAccount{owner: "John", numAgency: 123, numAccount: 123456, funds: 500}
	accounts["654321"] = CheckingAccount{owner: "Mary", numAgency: 321, numAccount: 654321, funds: 1500}
	accounts["456789"] = CheckingAccount{owner: "Peter", numAgency: 789, numAccount: 456789, funds: 2500}

	accountJohn := accounts["123456"]
	accountMary := accounts["654321"]
	accountPeter := accounts["456789"]

	accountJohn.WithdrawFunds(100)
	accountMary.DepositFunds(100)
	accountPeter.TransferFunds(100, &accountJohn)

	fmt.Println(accountJohn, accountMary, accountPeter)
}
