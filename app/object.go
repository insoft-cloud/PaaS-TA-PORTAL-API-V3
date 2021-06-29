package app

import (
	"time"
)

type CreateApp struct {
	Name      string `json:"name" validate:"required"`
	Lifecycle struct {
		Data struct {
			Buildpacks []string `json:"buildpacks"`
			Stack      string   `json:"stack"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"lifecycle"`
	EnvironmentVariables struct {
	} `json:"environment_variables"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata"`
	Relationships struct {
		Space struct {
			Data struct {
				GUID string `json:"guid" validate:"required"`
			} `json:"data"`
		} `json:"space"`
	} `json:"relationships"`
}

type App struct {
	CreatedAt time.Time `json:"created_at"`
	GUID      string    `json:"guid"`
	Lifecycle struct {
		Data struct {
			Buildpacks []string `json:"buildpacks"`
			Stack      string   `json:"stack"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"lifecycle"`
	Links struct {
		CurrentDroplet struct {
			Href string `json:"href"`
		} `json:"current_droplet"`
		DeployedRevisions struct {
			Href string `json:"href"`
		} `json:"deployed_revisions"`
		Droplets struct {
			Href string `json:"href"`
		} `json:"droplets"`
		EnvironmentVariables struct {
			Href string `json:"href"`
		} `json:"environment_variables"`
		Packages struct {
			Href string `json:"href"`
		} `json:"packages"`
		Processes struct {
			Href string `json:"href"`
		} `json:"processes"`
		Revisions struct {
			Href string `json:"href"`
		} `json:"revisions"`
		RouteMappings struct {
			Href string `json:"href"`
		} `json:"route_mappings"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Space struct {
			Href string `json:"href"`
		} `json:"space"`
		Start struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"start"`
		Stop struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"stop"`
		Tasks struct {
			Href string `json:"href"`
		} `json:"tasks"`
	} `json:"links"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata"`
	Name          string `json:"name"`
	Relationships struct {
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
	} `json:"relationships"`
	State     string    `json:"state"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AppList struct {
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
	Resources []App
}

type UpdateApp struct {
	Lifecycle struct {
		Data struct {
			Buildpacks []string `json:"buildpacks"`
			Stack      string   `json:"stack"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"lifecycle"`
	Metadata struct {
		Annotations struct {
		} `json:"annotations"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata"`
	Name string `json:"name" validate:"required"`
}

type AppDroplet struct {
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
