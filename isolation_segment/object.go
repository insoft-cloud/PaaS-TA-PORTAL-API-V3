package isolation_segments

import "time"

//*struct omitempty" post로 보낼 객체들만 선언
type IsolationSegment struct {
	GUID      string    `json:"guid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Organizations struct {
			Href string `json:"href"`
		} `json:"organizations"`
	} `json:"links"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata"`
}

type CreateIsolationSegment struct {
	Name     string `json:"name" validate:"required"`
	Metadata *struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata" validate:"required"`
}

type IsolationSegments struct {
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
		GUID      string    `json:"guid"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Links     struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Organizations struct {
				Href string `json:"href"`
			} `json:"organizations"`
		} `json:"links"`
		Metadata struct {
			Annotations struct {
			} `json:"annotations"`
			Labels struct {
			} `json:"labels"`
		} `json:"metadata"`
	} `json:"resources"`
}

type OrganizationsRelationship struct {
	Data []struct {
		GUID string `json:"guid"`
	} `json:"data"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Related struct {
			Href string `json:"href"`
		} `json:"related"`
	} `json:"links"`
}

type SpacesRelationship struct {
	Data []struct {
		GUID string `json:"guid"`
	} `json:"data"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type UpdateIsolationSegment struct {
	Name     string `json:"name,omitempty"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata,omitempty"`
}

type EntitleOrganizationsIsolationSegment struct {
	Data []struct {
		GUID string `json:"guid"`
	} `json:"data" validate:"required"`
}
