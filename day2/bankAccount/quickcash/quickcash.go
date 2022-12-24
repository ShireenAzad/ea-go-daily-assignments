package quickcash

type QuickCash struct {
	Account[] Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	var withDrawCash float64;
	var accountName string;
	loop:for _, account := range qc.Account {
		if(account.CanWithDraw(amount)){
			withDrawCash,accountName= amount, account.GetIdentifier();
			break loop
		}
    }
	if(len(accountName)==0){
		return withDrawCash,"In sufficient Funds";
	}
	return withDrawCash,accountName;
	
}
