package constants

type Currency string

const (
	ARS  Currency = "ARS"
	BRL  Currency = "BRL"
	CLP  Currency = "CLP"
	MXN  Currency = "MXN"
	COP  Currency = "COP"
	PEN  Currency = "PEN"
	UYU  Currency = "UYU"
	VES  Currency = "VES"
	MCN  Currency = "MCN"
	BTC  Currency = "BTC"
	USD  Currency = "USD"
	USDP Currency = "USDP"
	DCE  Currency = "DCE"
	ETH  Currency = "ETH"
	FDI  Currency = "FDI"
	CDB  Currency = "CDB"
)

func (currency_id Currency) String() string {
	return string(currency_id)
}

func (*Currency) GetValues() []string {
	return []string{
		ARS.String(),
		BRL.String(),
		CLP.String(),
		MXN.String(),
		COP.String(),
		PEN.String(),
		UYU.String(),
		VES.String(),
		MCN.String(),
		BTC.String(),
		USD.String(),
		USDP.String(),
		DCE.String(),
		ETH.String(),
		FDI.String(),
		CDB.String(),
	}
}
