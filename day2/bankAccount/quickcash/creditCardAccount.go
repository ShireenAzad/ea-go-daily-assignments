package quickcash
type CreditCardAccount struct{}
func(ca *CreditCardAccount) WithDraw(amount float64){
}
func(ca *CreditCardAccount) CanWithDraw(amount float64)bool{
	return true;
}
func(ca *CreditCardAccount) GetIdentifier() string{
	return "CREDIT_CARD_ACCOUNT";
}
type CreditCardAccountWithZeroBalance struct{}
func(ca *CreditCardAccountWithZeroBalance) WithDraw(amount float64){
}
func(ca *CreditCardAccountWithZeroBalance) CanWithDraw(amount float64)bool{
	return false;
}
func(ca *CreditCardAccountWithZeroBalance) GetIdentifier() string{
	return "ZERO_CREDIT_CARD_ACCOUNT";
}
