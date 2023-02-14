package types

import S "github.com/espinosajuanma/slingr-go"

const PLATFORM_TICKET = "support.platform.platformTickets"

type PlatformTicketPayload struct {
	Title       string `json:"title"`
	Project     string `json:"project"`
	Type        string `json:"type"`
	Priority    string `json:"priority"`
	Category    string `json:"category"`
	Account     string `json:"account"`
	Description string `json:"description"`
	UseTemplate bool   `json:"useTemplate"`
}

type PlatformTicket struct {
	Id      string `json:"id"`
	Label   string `json:"label"`
	Version int    `json:"version"`
	Number  int    `json:"number"`
	Project struct {
		Id                     string `json:"id"`
		Label                  string `json:"label"`
		Type                   string `json:"type"`
		VisibleToAllAssociates bool   `json:"visibleToAllAssociates"`
	} `json:"project"`
	UseTemplate bool `json:"useTemplate"`
	Template    struct {
		Id          string `json:"id"`
		Label       string `json:"label"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"template"`
	Title             string            `json:"title"`
	Type              string            `json:"type"`
	Priority          string            `json:"priority"`
	SupportAfterHours bool              `json:"supportAfterHours"`
	Category          string            `json:"category"`
	Account           S.RecordReference `json:"account"`
	App               S.RecordReference `json:"app"`
	Approvers         []struct {
		Approver       S.RecordReference `json:"approver"`
		ApprovalStatus string            `json:"approvalStatus"`
	} `json:"approvers"`
	Customers             []S.RecordReference `json:"customers"`
	Assignees             []S.RecordReference `json:"assignees"`
	Watchers              []S.RecordReference `json:"watchers"`
	RelatedPlatformIssues []struct {
		IssueNumber int  `json:"issueNumber"`
		Released    bool `json:"released"`
	} `json:"relatedPlatformIssues"`
	Status           string `json:"status"`
	Description      string `json:"description"`
	CalculatedStatus string `json:"calculatedStatus"`
	Notes            []struct {
		Note    string `json:"note"`
		AddedBy struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"addedBy"`
		TimeStamp int64 `json:"timeStamp"`
	} `json:"notes"`
	Files []struct {
		Description string            `json:"description"`
		AddedBy     S.RecordReference `json:"addedBy"`
		TimeStamp   int64             `json:"timeStamp"`
		Link        string            `json:"link"`
	} `json:"files"`
	CreatedBy S.RecordReference   `json:"createdBy"`
	Created   int64               `json:"created"`
	Modified  int64               `json:"modified"`
	Rank      string              `json:"rank"`
	People    []S.RecordReference `json:"people"`
}

type ManyPlatformTickets struct {
	Offset string           `json:"offset"`
	Total  int              `json:"total"`
	Items  []PlatformTicket `json:"items"`
}
