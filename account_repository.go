package kusopr_pay

type (
	AccountRepisotry struct {
		ds *DataSource
	}
)

func (rcv *AccountRepisotry) FindBy(id AccountID) (*Account, error) {
	var result *Account
	err := rcv.ds.ExecQuery(result, id)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (rcv *AccountRepisotry) AddAmount(id AccountID, amt Amount) error {
	return rcv.ds.ExecCommand(id, amt)
}

func (rcv *AccountRepisotry) MinusAmount(id AccountID, amt Amount) error {
	return rcv.ds.ExecCommand(id, amt)
}
