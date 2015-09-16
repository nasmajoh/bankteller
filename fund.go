package funding

// Fund holding a balance
type Fund struct {
	//unexported ie private variable
	balance int
}

//Newfund function returning a pointer to a fund.
func Newfund(initialBalance int) *Fund {
	return &Fund{
		balance: initialBalance,
	}
}

// Balance method returning fund balance.
func (f *Fund) Balance() int {
	return f.balance
}

//Withdraw method withdraws given amount from fund.
func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
