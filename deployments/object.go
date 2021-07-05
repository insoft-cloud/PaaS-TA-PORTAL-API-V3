package deployments

import "time"

type Deployment struct {
	GUID   string `json:"guid"`
	Status struct {
		Value   string `json:"value"`
		Reason  string `json:"reason"`
		Details struct {
			LastSuccessfulHealthcheck time.Time `json:"last_successful_healthcheck"`
		} `json:"details"`
	} `json:"status"`
	Strategy string `json:"strategy"`
	Droplet  struct {
		GUID string `json:"guid"`
	} `json:"droplet"`
	PreviousDroplet struct {
		GUID string `json:"guid"`
	} `json:"previous_droplet"`
	NewProcesses []struct {
		GUID string `json:"guid"`
		Type string `json:"type"`
	} `json:"new_processes"`
	Revision struct {
		GUID    string `json:"guid"`
		Version int    `json:"version"`
	} `json:"revision"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metadata  struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
		Cancel struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"cancel"`
	} `json:"links"`
}

type DeploymentList struct {
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
	Resources []Deployment `json:"resources"`
}

type CreateDeployment struct {
	Droplet  *struct{} `json:"droplet,omitempty"`
	Revision *struct{} `json:"droplet,omitempty"`
	Strategy string    `json:"strategy,omitempty"`
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid"  validate:"required"`
			} `json:"data"`
		} `json:"app"`
	} `json:"relationships"`
}

type UpdateDeployment struct {
	Metadata *struct {
		Labels *struct {
		} `json:"labels,omitempty"`
		Annotations *struct {
		} `json:"annotations,omitempty"`
	} `json:"metadata,omitempty"`
}
