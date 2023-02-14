package types

import S "github.com/espinosajuanma/slingr-go"

const FRONTEND_SOW_ENTITY = "frontendBilling.sow"

type ManyFrontendSOW struct {
	Total  int           `json:"total"`
	Offset string        `json:"offset"`
	Items  []FrontendSOW `json:"items"`
}

type FrontendSOW struct {
	Id                    string            `json:"id"`
	Version               int               `json:"version"`
	Label                 string            `json:"label"`
	Entity                S.EntityReference `json:"entity"`
	Status                string            `json:"status"`
	OverridePrices        S.RecordReference `json:"overridePrices"`
	OverrideNotifications struct {
		EnableOverrideNotifications bool   `json:"enableOverrideNotifications"`
		Id                          string `json:"id"`
		Label                       string `json:"label"`
	} `json:"overrideNotifications"`
	People []struct {
		User              S.RecordReference `json:"user"`
		Dedication        float64           `json:"dedication"`
		DefaultService    S.RecordReference `json:"defaultService"`
		DefaultSowService S.RecordReference `json:"defaultSowService"`
		Id                string            `json:"id"`
		Label             string            `json:"label"`
	} `json:"people"`
}
