package model

type TypeField int

const (
	Ip TypeField = iota
	Timestamp
	Domain
	Blacklisted
	EventType
)

var TypeFieldsValues = [...]string{"ip", "timestamp", "domain", "blacklisted", "event_type"}

func (field TypeField) String() string {
	return TypeFieldsValues[field]
}

type Document struct {
	Ip            string `json:"ip"`
	Timestamp     string `json:"timestamp"`
	Domain        string `json:"domain"`
	IsBlacklisted bool   `json:"blacklisted"`
	EventType     string `json:"event_type"`
}

type Documents []Document
