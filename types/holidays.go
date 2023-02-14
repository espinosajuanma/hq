package types

const HOLIDAY_ENTITY = "management.holidays"

type ManyHolidays struct {
	Total  int       `json:"total"`
	Offset string    `json:"offset"`
	Items  []Holiday `json:"items"`
}

type Holiday struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
	Label   string `json:"label"`
	Entity  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"entity"`
	Country struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"country"`
	Day    string `json:"day"`
	Title  string `json:"title"`
	Ignore bool   `json:"ignore"`
}
