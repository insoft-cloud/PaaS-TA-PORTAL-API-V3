package space_features

type SpaceFeature struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

type SpaceFeatureList struct {
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
	Resources []SpaceFeature `json:"resources"`
}
