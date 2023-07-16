package ui

import (
	"fmt"

	"github.com/fatih/color"

	"poogolang.com/m/v2/accounts"
)

func ShowMenu() {
	color.New(color.FgBlue).Println("** Bem vindo aoo Poobank **")
	fmt.Println("1 - Listar contas")
	fmt.Println("0 - Sair")
}

func ReadOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func ShowListAccountsMenu(
	checkingAccounts *[]accounts.CheckingAccount,
	savingsAccounts *[]accounts.SavingsAccount,
) {
	fmt.Println("1 - Contas correntes")
	fmt.Println("2 - Contas poupanca")
	fmt.Println("0 - Voltar")

	subOption := -1

	for subOption != 0 {
		fmt.Println("Escolha uma opcao: ")
		fmt.Scan(&subOption)

		switch subOption {
		case 1:
			ShowCheckingAccounts(checkingAccounts)
		case 2:
			ShowSavingsAccounts(savingsAccounts)
		case 0:
			fmt.Println("Voltando...")
		default:
			fmt.Println("Opcao invalida")
		}
	}
}

func ShowCheckingAccounts(checkingAccounts *[]accounts.CheckingAccount) {
	fmt.Println("Contas correntes")
	for _, checkingAccount := range *checkingAccounts {
		fmt.Println("*** Conta corrente ***")
		color.New(color.FgHiMagenta).Println("Genero: ", checkingAccount.Holder.Gender)
		color.New(color.FgHiMagenta).Println("Nome: ", checkingAccount.Holder.Name)
		color.New(color.FgHiMagenta).Println("CPF: ", checkingAccount.Holder.Cpf)
		color.New(color.FgHiMagenta).Println("Profissao: ", checkingAccount.Holder.Profission)
		color.New(color.FgHiMagenta).Println("Email: ", checkingAccount.Holder.Email)
		color.New(color.FgHiMagenta).Println("Saldo: ", checkingAccount.Funds)
		color.New(color.FgHiMagenta).Println("Numero da agencia: ", checkingAccount.NumAgency)
		color.New(color.FgHiMagenta).Println("Numero da conta: ", checkingAccount.NumAccount)
		fmt.Println("**********************")
	}
}

func ShowSavingsAccounts(savingsAccounts *[]accounts.SavingsAccount) {
	fmt.Println("Contas poupança")
	for _, savingsAccounts := range *savingsAccounts {
		fmt.Println("*** Conta poupança ***")
		color.New(color.FgHiYellow).Println("Genero: ", savingsAccounts.Holder.Gender)
		color.New(color.FgHiYellow).Println("Nome: ", savingsAccounts.Holder.Name)
		color.New(color.FgHiYellow).Println("CPF: ", savingsAccounts.Holder.Cpf)
		color.New(color.FgHiYellow).Println("Profissao: ", savingsAccounts.Holder.Profission)
		color.New(color.FgHiYellow).Println("Email: ", savingsAccounts.Holder.Email)
		color.New(color.FgHiYellow).Println("Saldo: ", savingsAccounts.Funds)
		color.New(color.FgHiYellow).Println("Numero da agencia: ", savingsAccounts.NumAgency)
		color.New(color.FgHiYellow).Println("Numero da conta: ", savingsAccounts.NumAccount)
		color.New(color.FgHiYellow).Println("Operação: ", savingsAccounts.Operation)
		fmt.Println("**********************")
	}
}