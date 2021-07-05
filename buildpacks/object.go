package buildpacks

import "time"

type BuildPack struct {
	GUID      string      `json:"guid"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
	State     string      `json:"state"`
	Filename  interface{} `json:"filename"`
	Stack     string      `json:"stack"`
	Position  int         `json:"position"`
	Enabled   bool        `json:"enabled"`
	Locked    bool        `json:"locked"`
	Metadata  struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Upload struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"upload"`
	} `json:"links"`
}

type BuildPackList struct {
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
	Resources []BuildPack `json:"resources"`
}

type CreateBuildPack struct {
	Name     string `json:"name" validate:"required"`
	Position int    `json:"position,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
	Locked   bool   `json:"locked,omitempty"`
	Stack    string `json:"stack,omitempty"`
	Metadata *struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata,omitempty"`
}

type UpdateBuildPack struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
	Enabled  bool   `json:"enabled"`
	Locked   bool   `json:"locked"`
	Stack    string `json:"stack"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata"`
}
