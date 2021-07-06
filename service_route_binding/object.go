package service_route_binding

import "time"

type ServiceRouteBinding struct {
	GUID            string    `json:"guid"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	RouteServiceURL string    `json:"route_service_url"`
	LastOperation   struct {
		Type        string    `json:"type"`
		State       string    `json:"state"`
		Description string    `json:"description"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedAt   time.Time `json:"created_at"`
	} `json:"last_operation"`
	Metadata struct {
		Annotations struct {
			Foo string `json:"foo"`
		} `json:"annotations"`
		Labels struct {
			Baz string `json:"baz"`
		} `json:"labels"`
	} `json:"metadata"`
	Relationships struct {
		ServiceInstance struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_instance"`
		Route struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"route"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		ServiceInstance struct {
			Href string `json:"href"`
		} `json:"service_instance"`
		Route struct {
			Href string `json:"href"`
		} `json:"route"`
		Parameters struct {
			Href string `json:"href"`
		} `json:"parameters"`
	} `json:"links"`
}

type ServiceRouteBindingList struct {
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
	Resources []ServiceRouteBinding `json:"resources"`
}

type CreateServiceRouteBinding struct {
	Metadata *struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata,omitempty"`
	Relationships struct {
		Route struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"route" validate:"required"`
		ServiceInstance struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_instance" validate:"required"`
	} `json:"relationships"`
	Parameters *struct {
	} `json:"parameters,omitempty"`
}

type UpdateServiceRouteBinding struct {
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata",omitempty`
}
