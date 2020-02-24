package kusopr_pay

type (
	TransferUsecase struct {
		accountRepo *AccountRepisotry
	}
)

func(rcv *TransferUsecase) Transfer(srcID AccountID, amt Amount, targetID AccountID) error {
	srcAccount, err := rcv.accountRepo.FindBy(srcID)
	if err != nil  {
		return err
	}

	targetAccount, err := rcv.accountRepo.FindBy(targetID)
	if err != nil  {
		return err
	}

	id, amnt, tgt, err := srcAccount.Transfer(amt, targetAccount)
	if err != nil {
		return err
	}

	err = rcv.accountRepo.AddAmount(tgt, amnt)
	if err == nil {
		return err
	}

	err = rcv.accountRepo.MinusAmount(id, amnt)
	if err == nil {
		return err
	}

	return nil
}
