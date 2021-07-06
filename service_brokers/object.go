package service_brokers

import "time"

type ServiceBroker struct {
	GUID          string    `json:"guid"`
	Name          string    `json:"name"`
	URL           string    `json:"url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Relationships struct {
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
	} `json:"relationships"`
	Metadata struct {
		Labels struct {
			Type string `json:"type"`
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		ServiceOfferings struct {
			Href string `json:"href"`
		} `json:"service_offerings"`
		Space struct {
			Href string `json:"href"`
		} `json:"space"`
	} `json:"links"`
}

type ServiceBrokerList struct {
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
	Resources []ServiceBroker `json:"resources"`
}

type CreateServiceBroker struct {
	Name           string `json:"name" validate:"required"`
	URL            string `json:"url" validate:"required"`
	Authentication struct {
		Type        string `json:"type" validate:"required"`
		Credentials struct {
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		} `json:"credentials" validate:"required"`
	} `json:"authentication" validate:"required"`
	Relationships *struct {
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
	} `json:"relationships,omitempty"`
	Metadata *struct {
		Labels struct {
			Type string `json:"type"`
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type UpdateServiceBroker struct {
	Name           string `json:"name,omitempty"`
	URL            string `json:"url,omitempty"`
	Authentication *struct {
		Type        string `json:"type"`
		Credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"credentials"`
	} `json:"authentication,omitempty"`
	Metadata *struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}
