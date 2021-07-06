package audit_events

import "time"

type AuditEvent struct {
	GUID      string    `json:"guid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      string    `json:"type"`
	Actor     struct {
		GUID string `json:"guid"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"actor"`
	Target struct {
		GUID string `json:"guid"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"target"`
	Data struct {
		Request struct {
			Recursive bool `json:"recursive"`
		} `json:"request"`
	} `json:"data"`
	Space struct {
		GUID string `json:"guid"`
	} `json:"space"`
	Organization struct {
		GUID string `json:"guid"`
	} `json:"organization"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type AuditEventList struct {
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
	Resources []AuditEvent `json:"resources"`
}
