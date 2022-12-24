package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}
func TestFailureWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	expectedErrorMessage :=acc.Withdraw(600)
	assert.EqualError(t,expectedErrorMessage,"In sufficient funds . You are totalAmount:500.000000.You are lacking by -100.000000")

}
