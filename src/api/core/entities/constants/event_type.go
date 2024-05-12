package constants

type EventType string

const (
	PaymentCreated          EventType = "payment.created"
	PaymentUpdated          EventType = "payment.updated"
	ApplicationDeauthorized EventType = "application.deauthorized"
	ApplicationAuthorized   EventType = "application.authorized"
	CreatedUpdated          EventType = "created - updated"
	StateFinished           EventType = "state_FINISHED"
	StateCanceled           EventType = "state_CANCELED"
	StateError              EventType = "state_ERROR"
	DeliveryUpdated         EventType = "delivery.updated"
	CaseCreated             EventType = "case_created"
	Updated                 EventType = "updated"
)

func (event_type EventType) String() string {
	return string(event_type)
}

func (*EventType) GetValues() []string {
	return []string{
		PaymentCreated.String(),
		PaymentUpdated.String(),
		ApplicationDeauthorized.String(),
		ApplicationAuthorized.String(),
		CreatedUpdated.String(),
		StateFinished.String(),
		StateCanceled.String(),
		StateError.String(),
		DeliveryUpdated.String(),
		CaseCreated.String(),
		Updated.String(),
	}
}

func (event_type EventType) IsPaymentCreated() bool {
	return event_type == PaymentCreated
}
