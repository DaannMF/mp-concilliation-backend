package constants

type PaymentStatusDetail string

const (
	Accredited                      PaymentStatusDetail = "Accredited"
	PendingContingency              PaymentStatusDetail = "pending_contingency"
	PendingReviewManual             PaymentStatusDetail = "pending_review_manual"
	CCRejectedBadFilledDate         PaymentStatusDetail = "cc_rejected_bad_filled_date"
	CCRejectedBadFilledOther        PaymentStatusDetail = "cc_rejected_bad_filled_other"
	CCRejectedBadFilledSecurityCode PaymentStatusDetail = "cc_rejected_bad_filled_security_code"
	CCRejectedBlacklist             PaymentStatusDetail = "cc_rejected_blacklist"
	CCRejectedCallForAuthorize      PaymentStatusDetail = "cc_rejected_call_for_authorize"
	CCRejectedCardDisabled          PaymentStatusDetail = "cc_rejected_card_disabled"
	CCRejectedDuplicatedPayment     PaymentStatusDetail = "cc_rejected_duplicated_payment"
	CCRejectedHighRisk              PaymentStatusDetail = "cc_rejected_high_risk"
	CCRejectedInsufficientAmount    PaymentStatusDetail = "cc_rejected_insufficient_amount"
	CCRejectedInvalidInstallments   PaymentStatusDetail = "cc_rejected_invalid_installments"
	CCRejectedMaxAttempts           PaymentStatusDetail = "cc_rejected_max_attempts"
	CCRejectedOtherReason           PaymentStatusDetail = "cc_rejected_other_reason"
)

func (payment_status_detail PaymentStatusDetail) String() string {
	return string(payment_status_detail)
}

func (*PaymentStatusDetail) GetValues() []string {
	return []string{
		Accredited.String(),
		PendingContingency.String(),
		PendingReviewManual.String(),
		CCRejectedBadFilledDate.String(),
		CCRejectedBadFilledOther.String(),
		CCRejectedBadFilledSecurityCode.String(),
		CCRejectedBlacklist.String(),
		CCRejectedCallForAuthorize.String(),
		CCRejectedCardDisabled.String(),
		CCRejectedDuplicatedPayment.String(),
		CCRejectedHighRisk.String(),
		CCRejectedInsufficientAmount.String(),
		CCRejectedInvalidInstallments.String(),
		CCRejectedMaxAttempts.String(),
		CCRejectedOtherReason.String(),
	}
}
