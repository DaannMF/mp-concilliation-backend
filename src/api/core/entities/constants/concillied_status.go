package constants

type ConcilliedStatus string

const (
	ConcilliedPending ConcilliedStatus = "PENDING"
	Concillied        ConcilliedStatus = "CONCILLIED"
)

func (status ConcilliedStatus) String() string {
	return string(status)
}

func (*ConcilliedStatus) GetValues() []string {
	return []string{
		ConcilliedPending.String(),
		Concillied.String(),
	}
}
