package accounts

import (
	"poogolang.com/m/v2/clients"
)


type CheckingAccount struct {
	Holder clients.Holder
	NumAgency int
	NumAccount int
	Funds float64
}

func (c *CheckingAccount) WithdrawFunds(value float64) bool {
	canWithdraw := value <= c.Funds && value > 0
	if canWithdraw {
		c.Funds -= value
		return true
	}
	return false
}

func (c *CheckingAccount) DepositFunds(value float64) bool {
	canDeposit := value > 0
	if canDeposit {
		c.Funds += value
		return	true
	}
	return false
}

func (c *CheckingAccount) TransferFunds(value float64, destinyAccount *CheckingAccount) bool {
	canTransfer := value <= c.Funds && destinyAccount != nil && value > 0
	if  canTransfer{
		c.Funds -= value
		destinyAccount.Funds += value
		return true
	}
	return false
}
