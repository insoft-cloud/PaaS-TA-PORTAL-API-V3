package roles

import "time"

type Role struct {
	GUID          string    `json:"guid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Type          string    `json:"type"`
	Relationships struct {
		User struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"user"`
		Organization struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organization"`
		Space struct {
			Data interface{} `json:"data"`
		} `json:"space"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		User struct {
			Href string `json:"href"`
		} `json:"user"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
	} `json:"links"`
}

type CreateRole struct {
	Type          string `json:"type" validation:"required"`
	Relationships *struct {
		User struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"user"`
		Organization struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organization"`
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
	} `json:"relationships"  validation:"required"`
}

type RoleList struct {
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
	Resources []Role `json:"resources"`
}
