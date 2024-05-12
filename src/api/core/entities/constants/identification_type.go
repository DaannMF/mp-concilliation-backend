package constants

type IdentificationType string

const (
	CPF  IdentificationType = "CPF"
	CNPJ IdentificationType = "CNPJ"
	CUIT IdentificationType = "CUIT"
	CUIL IdentificationType = "CUIL"
	DNI  IdentificationType = "DNI"
	CURP IdentificationType = "CURP"
	RFC  IdentificationType = "RFC"
	CC   IdentificationType = "CC"
	RUT  IdentificationType = "RUT"
	CI   IdentificationType = "CI"
)

func (identification_type IdentificationType) String() string {
	return string(identification_type)
}

func (*IdentificationType) GetValues() []string {
	return []string{
		CPF.String(),
		CNPJ.String(),
		CUIT.String(),
		CUIL.String(),
		DNI.String(),
		CURP.String(),
		RFC.String(),
		CC.String(),
		RUT.String(),
		CI.String(),
	}
}
