package feature_flags

import (
	"time"
)

type FeatureFlags struct {
	Name               string    `json:"name"`
	Enabled            bool      `json:"enabled"`
	UpdatedAt          time.Time `json:"updated_at"`
	CustomErrorMessage string    `json:"custom_error_message"`
	Links              struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type GetFeatureFlags struct {
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
	Resources []struct {
		Name               string    `json:"name"`
		Enabled            bool      `json:"enabled"`
		UpdatedAt          time.Time `json:"updated_at"`
		CustomErrorMessage string    `json:"custom_error_message"`
		Links              struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"resources"`
}

type UpdateFeatureFlags struct {
	Enabled            bool   `json:"enabled,omitempty"`
	CustomErrorMessage string `json:"custom_error_message,omitempty"`
}
