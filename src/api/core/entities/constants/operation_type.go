package constants

type OperationType string

const (
	Investment        OperationType = "investment"
	RegularPayment    OperationType = "regular_payment"
	MoneyTransfer     OperationType = "money_transfer"
	RecurringPayment  OperationType = "recurring_payment"
	AccountFund       OperationType = "account_fund"
	PaymentAddition   OperationType = "payment_addition"
	CellphoneRecharge OperationType = "cellphone_recharge"
	POSPayment        OperationType = "pos_payment"
	MoneyExchange     OperationType = "money_exchange"
)

func (operation_type OperationType) String() string {
	return string(operation_type)
}

func (*OperationType) GetValues() []string {
	return []string{
		Investment.String(),
		RegularPayment.String(),
		MoneyTransfer.String(),
		RecurringPayment.String(),
		AccountFund.String(),
		PaymentAddition.String(),
		CellphoneRecharge.String(),
		POSPayment.String(),
		MoneyExchange.String(),
	}
}
