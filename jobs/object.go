package jobs

import "time"

type Jobs struct {
	GUID      string    `json:"guid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Operation string    `json:"operation"`
	State     string    `json:"state"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
	Errors   []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
}
