package quickcash
type PaytmWallet struct{}
func(pw *PaytmWallet) WithDraw(amount float64){
}
func(pw *PaytmWallet) CanWithDraw(amount float64)bool{
	return true;
}
func(pw *PaytmWallet) GetIdentifier() string{
	return "PAYTM_WALLET_ACCOUNT";
}
type PaytmWalletWithZeroBalance struct{}
func(pw *PaytmWalletWithZeroBalance) WithDraw(amount float64){
}
func(pw *PaytmWalletWithZeroBalance) CanWithDraw(amount float64)bool{
	return false;
}
func(pw *PaytmWalletWithZeroBalance) GetIdentifier() string{
	return "ZERO_PAYTM_WALLET_ACCOUNT";
}
