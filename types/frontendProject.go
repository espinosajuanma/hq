package types

import S "github.com/espinosajuanma/slingr-go"

const FRONTEND_PROJECTS_ENTITY = "frontend.projects"

type ManyFrontendProjects struct {
	Total  int               `json:"total"`
	Offset string            `json:"offset"`
	Items  []FrontendProject `json:"items"`
}

type FrontendProject struct {
	Id                     string            `json:"id"`
	Version                int               `json:"version"`
	Label                  string            `json:"label"`
	Entity                 S.EntityReference `json:"entity"`
	Name                   string            `json:"name"`
	Description            string            `json:"description"`
	DescriptionSpanish     string            `json:"descriptionSpanish"`
	Type                   string            `json:"type"`
	Parent                 S.RecordReference `json:"parent"`
	VisibleToAllAssociates bool              `json:"visibleToAllAssociates"`
	RulesDocument          S.RecordReference `json:"rulesDocument"`
	Support                struct {
		DefaultAssignee S.RecordReference `json:"defaultAssignee"`
		SupportCode     string            `json:"supportCode"`
		SupportEmail    string            `json:"supportEmail"`
		Id              string            `json:"id"`
		Label           string            `json:"label"`
	} `json:"support"`
	People []struct {
		User  S.RecordReference   `json:"user"`
		Roles []S.RecordReference `json:"roles"`
		Sow   S.RecordReference   `json:"sow"`
		Id    string              `json:"id"`
		Label string              `json:"label"`
	} `json:"people"`
	RequiredCertifications []S.RecordReference `json:"requiredCertifications"`
	OpeningWorkingHour     string              `json:"openingWorkingHour"`
	EndWorkingHour         string              `json:"endWorkingHour"`
	Timezone               string              `json:"timezone"`
	DevelopmentCycle       struct {
		ReleaseDay     string `json:"releaseDay"`
		PrePlanningDay string `json:"prePlanningDay"`
		PlanningDay    string `json:"planningDay"`
		DemoDay        string `json:"demoDay"`
		Id             string `json:"id"`
		Label          string `json:"label"`
	} `json:"developmentCycle"`
	Status        string `json:"status"`
	Notifications []struct {
		Category string `json:"category"`
		Types    []struct {
			Type            S.RecordReference `json:"type"`
			ChannelsEnabled []string          `json:"channelsEnabled"`
			Id              string            `json:"id"`
			Label           string            `json:"label"`
		} `json:"types"`
		Id    string `json:"id"`
		Label string `json:"label"`
	} `json:"notifications"`
	Resources struct {
		PublicSlackChannel   S.RecordReference `json:"publicSlackChannel"`
		InternalSlackChannel S.RecordReference `json:"internalSlackChannel"`
		DocumentsFolder      string            `json:"documentsFolder"`
		Id                   string            `json:"id"`
		Label                string            `json:"label"`
	} `json:"resources"`
	CreatedBy             S.RecordReference   `json:"createdBy"`
	Created               int64               `json:"created"`
	Modified              int64               `json:"modified"`
	CanModifyTimeTracking []S.EntityReference `json:"canModifyTimeTracking"`
}
