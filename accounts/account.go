package accounts

import "poogolang.com/m/v2/clients"

type Account struct {
	Holder clients.Holder
	NumAgency, NumAccount int	
	Funds float64
	Email string	
}

type checkAccount interface {
	WithdrawFunds(value float64) bool
	DepositFunds(value float64) bool
	TransferFunds(value float64, destinyAccount *Account) bool
}


type SavingsAccount struct {
    Account
	Operation int
}

type CheckingAccount struct {
    Account
}

func WithdrawFunds(checkAccount checkAccount, value float64) bool {
	return checkAccount.WithdrawFunds(value)
}

func (c *Account) WithdrawFunds(value float64) bool {
	canWithdraw := value <= c.Funds && value > 0
	if canWithdraw {
		c.Funds -= value
		return true
	}
	return false
}

func DepositFunds(checkAccount checkAccount, value float64) bool {
	return checkAccount.DepositFunds(value)
}

func (c *Account) DepositFunds(value float64) bool {
	canDeposit := value > 0
	if canDeposit {
		c.Funds += value
		return	true
	}
	return false
}

func TransferFunds(checkAccountOrigin checkAccount, value float64, destinyAccount *Account) bool {
	return checkAccountOrigin.TransferFunds(value, destinyAccount)
}

func (c *Account) TransferFunds(value float64, destinyAccount *Account) bool {
	canTransfer := value <= c.Funds && destinyAccount != nil && value > 0
	if  canTransfer{
		c.Funds -= value
		destinyAccount.Funds += value
		return true
	}
	return false
}
