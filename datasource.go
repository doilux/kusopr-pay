package kusopr_pay

type (
	DataSource struct {}
)

func (rcv *DataSource) ExecCommand(command ...interface{}) error {
	// do some command
	return nil
}

func (rcv *DataSource) ExecQuery(target interface{}, query ...interface{}) error {
	// do some command
	return nil
}