package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
	
)

func TestGetCashFromFakePrimaryAccount(t *testing.T) {

	fpa := &FakePrimaryAccount{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		Account: []Withdrawable{fpa,fsa},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromFakeSecondaryAccount(t *testing.T) {

	fpa := &FakePrimaryAccountWithZeroBalance{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		Account: []Withdrawable{fpa,fsa},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}
func TestGetCashFromCreditCardAccount(t *testing.T) {
	
	ca := &CreditCardAccount{}
	sa := &SavingsAccount{}
	pw := &PaytmWallet{}
	fqc := QuickCash{
		Account: []Withdrawable{ca,sa,pw},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, ca.GetIdentifier(), accType)
}

func TestGetCashFromSavingsAccount(t *testing.T) {

	ca := &CreditCardAccountWithZeroBalance{}
	sa := &SavingsAccount{}
	pw := &PaytmWallet{}
	fqc := QuickCash{
		Account: []Withdrawable{ca,sa,pw},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, sa.GetIdentifier(), accType)
}
func TestGetCashFromPaytmWallet(t *testing.T) {

	ca := &CreditCardAccountWithZeroBalance{}
	sa := &SavingsAccountWithZeroBalance{}
	pw := &PaytmWallet{}
	fqc := QuickCash{
		Account: []Withdrawable{ca,sa,pw},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, pw.GetIdentifier(), accType)
}
func TestGetNoCashIfNoAmountInAnyAccount(t *testing.T) {

	ca := &CreditCardAccountWithZeroBalance{}
	sa := &SavingsAccountWithZeroBalance{}
	pw := &PaytmWalletWithZeroBalance{}
	fqc := QuickCash{
		Account: []Withdrawable{ca,sa,pw},
	}
	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(0), amt)
	assert.Equal(t,"In sufficient Funds", accType)
}