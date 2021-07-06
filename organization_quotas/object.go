package organization_quotas

import "time"

type OrganizationQuota struct {
	GUID      string    `json:"guid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Apps      struct {
		TotalMemoryInMb      int `json:"total_memory_in_mb"`
		PerProcessMemoryInMb int `json:"per_process_memory_in_mb"`
		TotalInstances       int `json:"total_instances"`
		PerAppTasks          int `json:"per_app_tasks"`
	} `json:"apps"`
	Services struct {
		PaidServicesAllowed   bool `json:"paid_services_allowed"`
		TotalServiceInstances int  `json:"total_service_instances"`
		TotalServiceKeys      int  `json:"total_service_keys"`
	} `json:"services"`
	Routes struct {
		TotalRoutes        int `json:"total_routes"`
		TotalReservedPorts int `json:"total_reserved_ports"`
	} `json:"routes"`
	Domains struct {
		TotalDomains int `json:"total_domains"`
	} `json:"domains"`
	Relationships struct {
		Organizations struct {
			Data []struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organizations"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type CreateOrganizationQuotas struct {
	Name string `json:"name" validate:"required"`
	Apps struct {
		TotalMemoryInMb      int `json:"total_memory_in_mb"`
		PerProcessMemoryInMb int `json:"per_process_memory_in_mb"`
		TotalInstances       int `json:"total_instances"`
		PerAppTasks          int `json:"per_app_tasks"`
	} `json:"apps,omitempty"`
	Services struct {
		PaidServicesAllowed   bool `json:"paid_services_allowed"`
		TotalServiceInstances int  `json:"total_service_instances"`
		TotalServiceKeys      int  `json:"total_service_keys"`
	} `json:"services,omitempty"`
	Routes struct {
		TotalRoutes        int `json:"total_routes"`
		TotalReservedPorts int `json:"total_reserved_ports"`
	} `json:"routes,omitempty"`
	Domains struct {
		TotalDomains int `json:"total_domains"`
	} `json:"domains,omitempty"`
	Relationships struct {
		Organizations struct {
			Data []struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organizations"`
	} `json:"relationships,omitempty"`
}

type OrganizationQuotasList struct {
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
	Resources []OrganizationQuota `json:"resources"`
}

type ApplyOrganizationQuotas struct {
	Data []struct {
		GUID string `json:"guid" validate:"required"`
	} `json:"data"`
}

type UpdateOrganizationQuota struct {
	Name string `json:"name"`
	Apps struct {
		TotalMemoryInMb      int `json:"total_memory_in_mb"`
		PerProcessMemoryInMb int `json:"per_process_memory_in_mb"`
		TotalInstances       int `json:"total_instances"`
		PerAppTasks          int `json:"per_app_tasks"`
	} `json:"apps"`
	Services struct {
		PaidServicesAllowed   bool `json:"paid_services_allowed"`
		TotalServiceInstances int  `json:"total_service_instances"`
		TotalServiceKeys      int  `json:"total_service_keys"`
	} `json:"services"`
	Routes struct {
		TotalRoutes        int `json:"total_routes"`
		TotalReservedPorts int `json:"total_reserved_ports"`
	} `json:"routes"`
	Domains struct {
		TotalPrivateDomains int `json:"total_private_domains"`
	} `json:"domains"`
}
