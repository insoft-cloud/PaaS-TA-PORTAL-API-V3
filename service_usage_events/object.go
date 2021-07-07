package service_usage_events

import "time"

type ServiceUsageEvent struct {
	GUID      string    `json:"guid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	State     string    `json:"state"`
	Space     struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"space"`
	Organization struct {
		GUID string `json:"guid"`
	} `json:"organization"`
	ServiceInstance struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"service_instance"`
	ServicePlan struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"service_plan"`
	ServiceOffering struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"service_offering"`
	ServiceBroker struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"service_broker"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type ServiceUsageEventList struct {
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
	Resources []ServiceUsageEvent `json:"resources"`
}
