package model

// Fact - данные которые отправляются на сервер.
type Fact struct {
	PeriodStart         string
	PeriodEnd           string
	PeriodKey           string
	IndicatorToMoID     string
	IndicatorToMoFactID string
	Value               string
	FactTime            string
	IsPlan              string
	AuthUserID          string
	Comment             string
}
