package app_usage_events

import "time"

type AppUsageEvent struct {
	GUID      string    `json:"guid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	State     struct {
		Current  string `json:"current"`
		Previous string `json:"previous"`
	} `json:"state"`
	App struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"app"`
	Process struct {
		GUID string `json:"guid"`
		Type string `json:"type"`
	} `json:"process"`
	Space struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"space"`
	Organization struct {
		GUID string `json:"guid"`
	} `json:"organization"`
	Buildpack struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"buildpack"`
	Task struct {
		GUID string `json:"guid"`
		Name string `json:"name"`
	} `json:"task"`
	MemoryInMbPerInstance struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
	} `json:"memory_in_mb_per_instance"`
	InstanceCount struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
	} `json:"instance_count"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type AppUsageEventList struct {
	Pagination struct {
		TotalResults int `json:"total_results"`
		TotalPages   int `json:"total_pages"`
		First        struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"pagination"`
	Resources []AppUsageEvent `json:"resources"`
}
