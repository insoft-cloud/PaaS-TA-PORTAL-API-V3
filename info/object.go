package info

type PlatformInfo struct {
	Build      string `json:"build"`
	CliVersion struct {
		Minimum     string `json:"minimum"`
		Recommended string `json:"recommended"`
	} `json:"cli_version"`
	Custom struct {
		Arbitrary string `json:"arbitrary"`
	} `json:"custom"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Version     int    `json:"version"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Support struct {
			Href string `json:"href"`
		} `json:"support"`
	} `json:"links"`
}

type PlatformUsageSummary struct {
	UsageSummary struct {
		StartedInstances int `json:"started_instances"`
		MemoryInMb       int `json:"memory_in_mb"`
	} `json:"usage_summary"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}
