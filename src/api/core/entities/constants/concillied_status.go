package constants

import "strings"

type ConcilliedStatus string

const (
	ConcilliedPending       ConcilliedStatus = "PENDING"
	Concillied              ConcilliedStatus = "CONCILLIED"
	UnknownConcilliedStatus ConcilliedStatus = ""
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

func ParseConcilliedStatus(value string) ConcilliedStatus {
	concilliedStatus := GetEnumValueFromString(new(ConcilliedStatus), value)

	if concilliedStatus != nil {
		return ConcilliedStatus(strings.ToUpper(*concilliedStatus))
	}

	return UnknownConcilliedStatus
}
