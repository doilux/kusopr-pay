package kusopr_pay

import "fmt"

type (
	Account struct {
		id     AccountID
		userID userID
		amount Amount
	}
)

func (rcv *Account) GetID() AccountID {
	if rcv == nil {
		return AccountID(0)
	}

	return rcv.id
}

func (rcv *Account) Transfer(amt Amount, targetAccount *Account) (AccountID, Amount, AccountID, error) {
	var er error

	if rcv == nil {
		er = fmt.Errorf("rcv is nil")
	} else if targetAccount == nil {
		er = fmt.Errorf("target is nil")
	}

	transferedSourceAccount, err := rcv.minusAmount(amt)
	if err != nil {
		er = err
	}

	return transferedSourceAccount.GetID(), amt, targetAccount.GetID(), er

}

func (rcv *Account) minusAmount(amt Amount) (*Account, error) {
	if rcv == nil {
		return nil, fmt.Errorf("rcv is nil")
	}

	remain := rcv.amount - amt
	if remain < 0 {
		return nil, fmt.Errorf("insufficient amount")
	}
	return &Account{
		id:     rcv.id,
		userID: rcv.userID,
		amount: remain,
	}, nil
}
