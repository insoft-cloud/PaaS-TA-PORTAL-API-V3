package service_credential_bindings

import "time"

type ServiceCredentialBinding struct {
	GUID          string    `json:"guid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	LastOperation struct {
		Type      string    `json:"type"`
		State     string    `json:"state"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"last_operation"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata"`
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app"`
		ServiceInstance struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_instance"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Details struct {
			Href string `json:"href"`
		} `json:"details"`
		Parameters struct {
			Href string `json:"href"`
		} `json:"parameters"`
		ServiceInstance struct {
			Href string `json:"href"`
		} `json:"service_instance"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
	} `json:"links"`
}

type ServiceCredentialBindingList struct {
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
	Resources []ServiceCredentialBinding `json:"resources"`
}

type CreateServiceCredentialBinding struct {
	Type          string `json:"type" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Relationships struct {
		ServiceInstance struct {
			Data struct {
				GUID string `json:"guid" validate:"required"`
			} `json:"data" validate:"required"`
		} `json:"service_instance" validate:"required"`
		App *struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app,omitempty"`
	} `json:"relationships" validate:"required"`
	Parameters *struct {
	} `json:"parameters,omitempty"`
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type UpdateServiceCredentialBinding struct {
	Metadata *struct {
		Labels *struct {
		} `json:"labels,omitempty"`
		Annotations *struct {
		} `json:"annotations,omitempty"`
	} `json:"metadata,omitempty"`
}
