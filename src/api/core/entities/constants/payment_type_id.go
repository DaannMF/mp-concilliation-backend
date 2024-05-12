package constants

type PaymentTypeID string

const (
	Ticket          PaymentTypeID = "ticket"
	BankTransfer    PaymentTypeID = "bank_transfer"
	ATM             PaymentTypeID = "atm"
	CreditCard      PaymentTypeID = "credit_card"
	DebitCard       PaymentTypeID = "debit_card"
	PrepaidCard     PaymentTypeID = "prepaid_card"
	DigitalCurrency PaymentTypeID = "digital_currency"
	DigitalWallet   PaymentTypeID = "digital_wallet"
	VoucherCard     PaymentTypeID = "voucher_card"
	CryptoTransfer  PaymentTypeID = "crypto_transfer"
)

func (payment_method_id PaymentTypeID) String() string {
	return string(payment_method_id)
}

func (*PaymentTypeID) GetValues() []string {
	return []string{
		Ticket.String(),
		BankTransfer.String(),
		ATM.String(),
		CreditCard.String(),
		DebitCard.String(),
		PrepaidCard.String(),
		DigitalCurrency.String(),
		DigitalWallet.String(),
		VoucherCard.String(),
		CryptoTransfer.String(),
	}
}
