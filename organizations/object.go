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
	} `json:"metadata,omitempty"`
}

type OrganizationsList struct {
	Pagination struct {
		TotalResults int `json:"total_results"`
		TotalPages   int `json:"total_pages"`
		First        struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"pagination"`
	Resources []Organizations `json:"resources"`
}

type UpdateOrganizations struct {
	Name      string `json:"name,omitempty"`
	Suspended bool   `json:"suspended,omitempty"`
	Metadata  *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type DefaultIsolationSegmentOrganizations struct {
	Data *struct {
		GUID string `json:"guid" validate:"required"`
	} `json:"data"`
}

type GetDefaultDomain struct {
	GUID               string      `json:"guid"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	Name               string      `json:"name"`
	Internal           bool        `json:"internal"`
	RouterGroup        interface{} `json:"router_group"`
	SupportedProtocols []string    `json:"supported_protocols"`
	Metadata           struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Relationships struct {
		Organization struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organization"`
		SharedOrganizations struct {
			Data []struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"shared_organizations"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
		RouteReservations struct {
			Href string `json:"href"`
		} `json:"route_reservations"`
		SharedOrganizations struct {
			Href string `json:"href"`
		} `json:"shared_organizations"`
	} `json:"links"`
}

type GetUsageSummary struct {
	UsageSummary struct {
		StartedInstances int `json:"started_instances"`
		MemoryInMb       int `json:"memory_in_mb"`
	} `json:"usage_summary"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
	} `json:"links"`
}
