package service_plans

import "time"

type ServicePlan struct {
	GUID           string `json:"guid"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	VisibilityType string `json:"visibility_type"`
	Available      bool   `json:"available"`
	Free           bool   `json:"free"`
	Costs          []struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
		Unit     string  `json:"unit"`
	} `json:"costs"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	MaintenanceInfo struct {
		Version     string `json:"version"`
		Description string `json:"description"`
	} `json:"maintenance_info"`
	BrokerCatalog struct {
		ID       string `json:"id"`
		Metadata struct {
			CustomKey string `json:"custom-key"`
		} `json:"metadata"`
		MaximumPollingDuration interface{} `json:"maximum_polling_duration"`
		Features               struct {
			PlanUpdateable bool `json:"plan_updateable"`
			Bindable       bool `json:"bindable"`
		} `json:"features"`
	} `json:"broker_catalog"`
	Schemas struct {
		ServiceInstance struct {
			Create struct {
				Parameters struct {
					Schema     string `json:"$schema"`
					Type       string `json:"type"`
					Properties struct {
						BillingAccount struct {
							Description string `json:"description"`
							Type        string `json:"type"`
						} `json:"billing-account"`
					} `json:"properties"`
				} `json:"parameters"`
			} `json:"create"`
			Update struct {
				Parameters struct {
				} `json:"parameters"`
			} `json:"update"`
		} `json:"service_instance"`
		ServiceBinding struct {
			Create struct {
				Parameters struct {
				} `json:"parameters"`
			} `json:"create"`
		} `json:"service_binding"`
	} `json:"schemas"`
	Relationships struct {
		ServiceOffering struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_offering"`
	} `json:"relationships"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		ServiceOffering struct {
			Href string `json:"href"`
		} `json:"service_offering"`
		Visibility struct {
			Href string `json:"href"`
		} `json:"visibility"`
	} `json:"links"`
}

type ServicePlanList struct {
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
	Resources []ServicePlan `json:"resources"`
}

type UpdatePlan struct {
	Metadata *struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}
