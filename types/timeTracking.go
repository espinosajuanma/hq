package types

import S "github.com/espinosajuanma/slingr-go"

const TIME_TRACKING_ENTITY = "frontendBilling.timeTracking"

type TimeTrackingPayload struct {
	Project         GenericRecord `json:"project"`
	Person          GenericRecord `json:"person"`
	Day             string        `json:"day"`
	TimeSpent       int64         `json:"timeSpent"`
	Service         GenericRecord `json:"service"`
	Sow             GenericRecord `json:"sow"`
	ServiceCategory []string      `json:"serviceCategory"`
	Billable        bool          `json:"billable"`
	Notes           string        `json:"notes"`
}

type TimeTracking struct {
	Id        string            `json:"id"`
	Label     string            `json:"label,omitempty`
	Entity    S.EntityReference `json:"label,omitempty`
	Project   S.RecordReference `json:"project"`
	Person    S.RecordReference `json:"person"`
	Day       string            `json:"day"`
	Notes     string            `json"notes"`
	TimeSpent int64             `json:"timeSpent"`
	Billable  bool              `json:"billable"`
}

type ManyTimeTracking struct {
	Total  int            `json:"total"`
	Items  []TimeTracking `json:"items"`
	Offset string         `json:"offset"`
}
