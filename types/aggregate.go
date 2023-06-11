package types

type Aggregate []struct {
	Match interface{} `json:"match,omitempty"`
	Group interface{} `json:"group,omitempty"`
}
