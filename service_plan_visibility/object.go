package service_plan_visibility

type ServicePlanVisibility struct {
	Type          string `json:"type"`
	Organizations []struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"organizations"`
}

type ServicePlanVisibilityList struct {
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
	Resources []ServicePlanVisibility `json:"resources"`
}

type UpdateServicePlanVisibility struct {
	Type          string `json:"type,omitempty"`
	Organizations []*struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"organizations,omitempty"`
}
