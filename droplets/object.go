package droplets

import (
	"time"
)

type CreateDroplet struct {
	Relationships *struct {
		App struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app" validate:"required"`
	} `json:"relationships"`
	ProcessTypes *struct {
		Rake string `json:"rake"`
		Web  string `json:"web"`
	} `json:"process_types,omitempty"`
}

//Get a droplet
type Droplet struct {
	GUID      string      `json:"guid"`
	State     string      `json:"state"`
	Error     interface{} `json:"error"`
	Lifecycle struct {
		Type string `json:"type"`
		Data struct {
		} `json:"data"`
	} `json:"lifecycle"`
	ExecutionMetadata string `json:"execution_metadata"`
	ProcessTypes      struct {
		Rake string `json:"rake"`
		Web  string `json:"web"`
	} `json:"process_types"`
	Checksum struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"checksum"`
	Buildpacks []struct {
		Name          string `json:"name"`
		DetectOutput  string `json:"detect_output"`
		Version       string `json:"version"`
		BuildpackName string `json:"buildpack_name"`
	} `json:"buildpacks"`
	Stack         string      `json:"stack"`
	Image         interface{} `json:"image"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Package struct {
			Href string `json:"href"`
		} `json:"package"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
		AssignCurrentDroplet struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"assign_current_droplet"`
		Download struct {
			Href string `json:"href"`
		} `json:"download"`
	} `json:"links"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
}

type DropletList struct {
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
	Resources []Droplet `json:"resources"`
}

type UpdateDroplet struct {
	Metadata struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata"`
}

type CopyDroplet struct {
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid" validate:"required"`
			} `json:"data"`
		} `json:"app"`
	} `json:"relationships"`
}
