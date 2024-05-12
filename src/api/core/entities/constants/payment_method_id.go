package constants

type PaymentMethodID string

const (
	Pix           PaymentMethodID = "Pix"
	DebinTransfer PaymentMethodID = "Debin_transfer"
	Ted           PaymentMethodID = "Ted"
	CVU           PaymentMethodID = "CVU"
	PSE           PaymentMethodID = "PSE"
)

func (payment_method_id PaymentMethodID) String() string {
	return string(payment_method_id)
}

func (*PaymentMethodID) GetValues() []string {
	return []string{
		Pix.String(),
		DebinTransfer.String(),
		Ted.String(),
		CVU.String(),
		PSE.String(),
	}
}
