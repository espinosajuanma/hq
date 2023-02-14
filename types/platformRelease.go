package types

const PLATFORM_RELEASE = "platform.releases"

type PlatformRelease struct {
	Id                  string `json:"id"`
	Label               string `json:"label"`
	Version             int    `json:"version"`
	Number              int    `json:"number"`
	StartDate           int64  `json:"startDate"`
	EndDate             int64  `json:"endDate"`
	ReleaseDate         int64  `json:"releaseDate"`
	ReleaseNotes        string `json:"releaseNotes"`
	Notified            bool   `json:"notified"`
	MailchimpCampaignID string `json:"mailchimpCampaignId"`
}

type ManyPlatformReleases struct {
	Offset string            `json:"offset"`
	Total  int               `json:"total"`
	Items  []PlatformRelease `json:"items"`
}
