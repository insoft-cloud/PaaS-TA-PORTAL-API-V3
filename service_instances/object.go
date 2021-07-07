package service_instances

import "time"

type ServiceInstance struct {
	GUID            string        `json:"guid"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Name            string        `json:"name"`
	Tags            []interface{} `json:"tags"`
	Type            string        `json:"type"`
	MaintenanceInfo struct {
		Version string `json:"version"`
	} `json:"maintenance_info"`
	UpgradeAvailable bool   `json:"upgrade_available"`
	DashboardURL     string `json:"dashboard_url"`
	LastOperation    struct {
		Type        string    `json:"type"`
		State       string    `json:"state"`
		Description string    `json:"description"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedAt   time.Time `json:"created_at"`
	} `json:"last_operation"`
	Relationships struct {
		ServicePlan struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_plan"`
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
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
		ServicePlan struct {
			Href string `json:"href"`
		} `json:"service_plan"`
		Space struct {
			Href string `json:"href"`
		} `json:"space"`
		Parameters struct {
			Href string `json:"href"`
		} `json:"parameters"`
		SharedSpaces struct {
			Href string `json:"href"`
		} `json:"shared_spaces"`
		ServiceCredentialBindings struct {
			Href string `json:"href"`
		} `json:"service_credential_bindings"`
		ServiceRouteBindings struct {
			Href string `json:"href"`
		} `json:"service_route_bindings"`
	} `json:"links"`
}

type ServiceInstanceList struct {
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
	Resources []ServiceInstance `json:"resources"`
}

type CreateServiceInstance struct {
	Type       string `json:"type" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Parameters *struct {
	} `json:"parameters,omitempty"`
	Tags     []string `json:"tags"`
	Metadata *struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata,omitempty"`
	Relationships struct {
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space" validate:"required"`
		ServicePlan struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_plan" validate:"required"`
	} `json:"relationships"`
}

type CreateServiceInstanceProvide struct {
	Type        string `json:"type" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Credentials *struct {
	} `json:"credentials,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	SyslogDrainURL  string   `json:"syslog_drain_url,omitempty"`
	RouteServiceURL string   `json:"route_service_url,omitempty"`
	Metadata        *struct {
		Annotations struct {
			Foo string `json:"foo"`
		} `json:"annotations"`
		Labels struct {
			Baz string `json:"baz"`
		} `json:"labels"`
	} `json:"metadata,omitempty"`
	Relationships struct {
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space" validate:"required"`
	} `json:"relationships"`
}

type UpdateServiceInstance struct {
	Name       string `json:"name,omitempty"`
	Parameters *struct {
	} `json:"parameters,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Relationships *struct {
		ServicePlan struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_plan"`
	} `json:"relationships,omitempty"`
	Metadata *struct {
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
	} `json:"metadata,omitempty"`
}

type SharedSpacesRelationship struct {
	Data []struct {
		GUID string `json:"guid"`
	} `json:"data"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type SharedServiceInstanceToOtherSpaces struct {
	Data []struct {
		GUID string `json:"guid"`
	} `json:"data" validate:"required"`
}

type GetUsageSummaryInSharedSpace struct {
	UsageSummary []struct {
		Space struct {
			GUID string `json:"guid"`
		} `json:"space"`
		BoundAppCount int `json:"bound_app_count"`
	} `json:"usage_summary"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		SharedSpaces struct {
			Href string `json:"href"`
		} `json:"shared_spaces"`
		ServiceInstance struct {
			Href string `json:"href"`
		} `json:"service_instance"`
	} `json:"links"`
}
