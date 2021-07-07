package environment_variable_groups

import "time"

type EnvironmentVariableGroup struct {
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Var       struct {
		Foo string `json:"foo"`
	} `json:"var"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type UpdateEnvironmentVariableGroup struct {
	Var struct {
		Debug string      `json:"DEBUG"`
		User  interface{} `json:"USER"`
	} `json:"var"`
}
