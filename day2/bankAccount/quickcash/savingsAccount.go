package quickcash
type SavingsAccount struct{}

func (sa *SavingsAccount) WithDraw(amount float64) {
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	return true
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}
type SavingsAccountWithZeroBalance struct{}

func (sa *SavingsAccountWithZeroBalance) WithDraw(amount float64) {
}

func (sa *SavingsAccountWithZeroBalance) CanWithDraw(amount float64) bool {
	return false
}

func (sa *SavingsAccountWithZeroBalance) GetIdentifier() string {
	return "ZERO_SAVINGS_ACCOUNT"
}