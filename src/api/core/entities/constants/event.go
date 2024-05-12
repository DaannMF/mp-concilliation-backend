package constants

type Event string

const (
	PaymentEvt                    Event = "payment"
	MPConnect                     Event = "mp-connect"
	SubscriptionPreapproval       Event = "subscription_preapproval"
	SubscriptionPreapprovalPlan   Event = "subscription_preapproval_plan"
	SubscriptionAuthorizedPayment Event = "subscription_authorized_payment"
	PointIntegrationWH            Event = "point_integration_wh"
	Delivery                      Event = "delivery"
	DeliveryCancellation          Event = "delivery_cancellation"
	TopicClaimsIntegrationWH      Event = "topic_claims_integration_wh"
)

func (event Event) String() string {
	return string(event)
}

func (*Event) GetValues() []string {
	return []string{
		PaymentEvt.String(),
		MPConnect.String(),
		SubscriptionPreapproval.String(),
		SubscriptionPreapprovalPlan.String(),
		SubscriptionAuthorizedPayment.String(),
		PointIntegrationWH.String(),
		Delivery.String(),
		DeliveryCancellation.String(),
		TopicClaimsIntegrationWH.String(),
	}
}

func (event Event) IsPayment() bool {
	return event == PaymentEvt
}
