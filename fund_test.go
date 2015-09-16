package funding

import "testing"

func BenchmarkFund(b *testing.B) {
	// Add as many dollars as we have iterations this run
	fund := Newfund(b.N)

	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance was't zero: ", fund.Balance())
	}
}
