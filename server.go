package funding

import "fmt"

//FundServer handling withdraws
type FundServer struct {
	Commands chan interface{}
	fund     Fund
}

//WithdrawCommand for server
type WithdrawCommand struct {
	Amount int
}

//BalanceCommand for server
type BalanceCommand struct {
	Response chan int
}

//NewFundServer creates new FundServer
func NewFundServer(initialBalance int) *FundServer {
	server := FundServer{
		//make() creates builtnis like channels, maps and slices
		Commands: make(chan interface{}),
		fund:     *Newfund(initialBalance),
	}
	go server.loop()
	return &server
}

func (s *FundServer) loop() {
	for command := range s.Commands {
		switch command.(type) {
		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}
