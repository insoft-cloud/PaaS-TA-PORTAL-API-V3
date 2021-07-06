package service_offerings

import "time"

type ServiceOffering struct {
	GUID             string        `json:"guid"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Available        bool          `json:"available"`
	Tags             []string      `json:"tags"`
	Requires         []interface{} `json:"requires"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	Shareable        bool          `json:"shareable"`
	DocumentationURL string        `json:"documentation_url"`
	BrokerCatalog    struct {
		ID       string `json:"id"`
		Metadata struct {
			Shareable bool `json:"shareable"`
		} `json:"metadata"`
		Features struct {
			PlanUpdateable       bool `json:"plan_updateable"`
			Bindable             bool `json:"bindable"`
			InstancesRetrievable bool `json:"instances_retrievable"`
			BindingsRetrievable  bool `json:"bindings_retrievable"`
			AllowContextUpdates  bool `json:"allow_context_updates"`
		} `json:"features"`
	} `json:"broker_catalog"`
	Relationships struct {
		ServiceBroker struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"service_broker"`
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
		ServicePlans struct {
			Href string `json:"href"`
		} `json:"service_plans"`
		ServiceBroker struct {
			Href string `json:"href"`
		} `json:"service_broker"`
	} `json:"links"`
}

type ServiceOfferingList struct {
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
	Resources []ServiceOffering `json:"resources"`
}

type UpdateServiceOffering struct {
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}
