package account

import (
	"fmt"
)

// TDD Bank Account app
type LowBalance struct{
	withDrawlAmount float64
	totalAmount float64
	withDrawFailureMessage string

}
func (lb LowBalance) Error()string{
	return fmt.Sprintf("%s . You are totalAmount:%f.You are lacking by %f",lb.withDrawFailureMessage,lb.totalAmount,(lb.totalAmount-lb.withDrawlAmount));
}
type Account struct {
	balance float64
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64)(error) {
	if(amount>acc.balance){
	return LowBalance{amount,acc.balance,"In sufficient funds"}
	}
	acc.balance -= amount
	return nil
}
