package constants

type PaymentStatus string

const (
	Pending     PaymentStatus = "pending"
	Approved    PaymentStatus = "approved"
	Authorized  PaymentStatus = "authorized"
	InProcess   PaymentStatus = "in_process"
	InMediation PaymentStatus = "in_mediation"
	Rejected    PaymentStatus = "rejected"
	Cancelled   PaymentStatus = "cancelled"
	Refunded    PaymentStatus = "refunded"
	ChargedBack PaymentStatus = "charged_back"
)

func (payment_status PaymentStatus) String() string {
	return string(payment_status)
}

func (*PaymentStatus) GetValues() []string {
	return []string{
		Pending.String(),
		Approved.String(),
		Authorized.String(),
		InProcess.String(),
		InMediation.String(),
		Rejected.String(),
		Cancelled.String(),
		Refunded.String(),
		ChargedBack.String(),
	}
}
