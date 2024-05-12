package constants

type PayerType string

const (
	Customer PayerType = "customer"
	Guest    PayerType = "guest"
)

func (payer_type PayerType) String() string {
	return string(payer_type)
}

func (*PayerType) GetValues() []string {
	return []string{
		Customer.String(),
		Guest.String(),
	}
}
