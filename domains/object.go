package domains

import (
	"time"
)

type CreateDomain struct {
	Name        string      `json:"name" validate:"required"`
	Internal    bool        `json:"internal,omitempty"`
	RouterGroup interface{} `json:"router_group,omitempty"`
	Metadata    *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
	Relationships *struct {
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
	} `json:"relationships,omitempty"`
}
type Domain struct {
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

type DomainList struct {
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
	Resources []Domain `json:"resources"`
}

type OrganizationDomainsList struct {
	Pagination struct {
		TotalResults int `json:"total_results"`
		TotalPages   int `json:"total_pages"`
		First        struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"pagination"`
	Resources []struct {
		GUID        string    `json:"guid"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Name        string    `json:"name"`
		Internal    bool      `json:"internal"`
		RouterGroup struct {
			GUID string `json:"guid"`
		} `json:"router_group"`
		SupportedProtocols []string `json:"supported_protocols"`
		Metadata           struct {
			Labels struct {
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
		} `json:"metadata"`
		Relationships struct {
			Organization struct {
				Data interface{} `json:"data"`
			} `json:"organization"`
			SharedOrganizations struct {
				Data []interface{} `json:"data"`
			} `json:"shared_organizations"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			RouteReservations struct {
				Href string `json:"href"`
			} `json:"route_reservations"`
			RouterGroup struct {
				Href string `json:"href"`
			} `json:"router_group"`
		} `json:"links"`
	} `json:"resources"`
}

type UpdateDomains struct {
	Metadata struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata"`
}

type ShareDomains struct {
	Data []struct {
		GUID string `json:"guid" validate:"required"`
	} `json:"data" `
}
