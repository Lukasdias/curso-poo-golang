package accounts

type CheckingAccount struct {
	owner string
	numAgency int
	numAccount int
	funds float64
}

func (c *CheckingAccount) WithdrawFunds(value float64) bool {
	canWithdraw := value <= c.funds && value > 0
	if canWithdraw {
		c.funds -= value
		return true
	}
	return false
}

func (c *CheckingAccount) DepositFunds(value float64) bool {
	canDeposit := value > 0
	if canDeposit {
		c.funds += value
		return	true
	}
	return false
}

func (c *CheckingAccount) TransferFunds(value float64, destinyAccount *CheckingAccount) bool {
	canTransfer := value <= c.funds && destinyAccount != nil && value > 0
	if  canTransfer{
		c.funds -= value
		destinyAccount.funds += value
		return true
	}
	return false
}
