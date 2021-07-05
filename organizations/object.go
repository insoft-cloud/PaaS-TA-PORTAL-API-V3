package organizations

import "time"

type Organizations struct {
	GUID          string    `json:"guid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Suspended     bool      `json:"suspended"`
	Relationships struct {
		Quota struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"quota"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Domains struct {
			Href string `json:"href"`
		} `json:"domains"`
		DefaultDomain struct {
			Href string `json:"href"`
		} `json:"default_domain"`
		Quota struct {
			Href string `json:"href"`
		} `json:"quota"`
	} `json:"links"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
}

type CreateOrganizations struct {
	Name      string `json:"name" validate:"required"`
	Suspended bool   `json:"suspended,omitempty"`
	Metadata  struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
}
