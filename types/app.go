package types

type CurrentUser struct {
	Localization struct {
		Timezone     string `json:"timezone"`
		TimezoneMode string `json:"timezoneMode"`
		Lang         string `json:"lang"`
	} `json:"localization"`
	LastName          string `json:"lastName"`
	IdentityProviders []struct {
		Id         string      `json:"id"`
		Name       string      `json:"name"`
		ExternalId interface{} `json:"externalId"`
		Label      string      `json:"label"`
	} `json:"identityProviders"`
	FullName string `json:"fullName"`
	Groups   []struct {
		Id       string `json:"id"`
		Primary  bool   `json:"primary"`
		External bool   `json:"external"`
		Name     string `json:"name"`
		Label    string `json:"label"`
	} `json:"groups"`
	Version                    int         `json:"version"`
	FirstName                  string      `json:"firstName"`
	PhoneNumber                interface{} `json:"phoneNumber"`
	TwoFactorAuthenticationKey interface{} `json:"twoFactorAuthenticationKey"`
	Permissions                struct {
		Developer       bool `json:"developer"`
		UsersManagement bool `json:"usersManagement"`
	} `json:"permissions"`
	TwoFactorAuthentication bool          `json:"twoFactorAuthentication"`
	Id                      string        `json:"id"`
	Integrations            []interface{} `json:"integrations"`
	Email                   string        `json:"email"`
	Status                  string        `json:"status"`
}
