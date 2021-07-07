package spaces

import "time"

type Space struct {
	GUID          string    `json:"guid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Relationships struct {
		Organization struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organization"`
		Quota struct {
			Data interface{} `json:"data"`
		} `json:"quota"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Features struct {
			Href string `json:"href"`
		} `json:"features"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
		ApplyManifest struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"apply_manifest"`
	} `json:"links"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
}

type SpaceList struct {
	Pagination struct {
		First struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next         interface{} `json:"next"`
		Previous     interface{} `json:"previous"`
		TotalPages   int         `json:"total_pages"`
		TotalResults int         `json:"total_results"`
	} `json:"pagination"`
	Resources []Space `json:"resources"`
}

type CreateSpace struct {
	Name          string `json:"name" validate:"required"`
	Relationships struct {
		Organization struct {
			Data struct {
				GUID string `json:"guid" validate:"required"`
			} `json:"data" validate:"required"`
		} `json:"organization" validate:"required"`
	} `json:"relationships" validate:"required"`
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type UpdateSpace struct {
	Name     string `json:"name,omitempty"`
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type IsolationSegment struct {
	Data struct {
		GUID string `json:"guid"`
	} `json:"data" validate:"required"`
}
