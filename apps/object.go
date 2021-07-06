package apps

import (
	"time"
)

type CreateApp struct {
	Name      string `json:"name" validate:"required"`
	Lifecycle *struct {
	} `json:"lifecycle,omitempty"`
	EnvironmentVariables *struct {
	} `json:"environment_variables,omitempty"`
	Metadata *struct {
		Annotations struct {
		} `json:"annotations,omitempty"`
		Labels struct {
		} `json:"labels"`
	} `json:"metadata,omitempty"`
	Relationships *struct {
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
	Resources []App `json:"resources"`
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

type AppDropletAssociation struct {
	Data struct {
		GUID string `json:"guid"`
	} `json:"data"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Related struct {
			Href string `json:"href"`
		} `json:"related"`
	} `json:"links"`
}

type AppEnv struct {
	StagingEnvJSON struct {
		GemCache string `json:"GEM_CACHE"`
	} `json:"staging_env_json"`
	RunningEnvJSON struct {
		HTTPProxy string `json:"HTTP_PROXY"`
	} `json:"running_env_json"`
	EnvironmentVariables struct {
		RailsEnv string `json:"RAILS_ENV"`
	} `json:"environment_variables"`
	SystemEnvJSON struct {
		VcapServices struct {
			Mysql []struct {
				Name         string   `json:"name"`
				BindingID    string   `json:"binding_id"`
				BindingName  string   `json:"binding_name"`
				InstanceID   string   `json:"instance_id"`
				InstanceName string   `json:"instance_name"`
				Label        string   `json:"label"`
				Tags         []string `json:"tags"`
				Plan         string   `json:"plan"`
				Credentials  struct {
					Username string `json:"username"`
					Password string `json:"password"`
				} `json:"credentials"`
				SyslogDrainURL string        `json:"syslog_drain_url"`
				VolumeMounts   []interface{} `json:"volume_mounts"`
				Provider       interface{}   `json:"provider"`
			} `json:"mysql"`
		} `json:"VCAP_SERVICES"`
	} `json:"system_env_json"`
	ApplicationEnvJSON struct {
		VcapApplication struct {
			Limits struct {
				Fds int `json:"fds"`
			} `json:"limits"`
			ApplicationName string      `json:"application_name"`
			ApplicationUris []string    `json:"application_uris"`
			Name            string      `json:"name"`
			SpaceName       string      `json:"space_name"`
			SpaceID         string      `json:"space_id"`
			Uris            []string    `json:"uris"`
			Users           interface{} `json:"users"`
		} `json:"VCAP_APPLICATION"`
	} `json:"application_env_json"`
}

type AppEnvVariable struct {
	Var struct {
	} `json:"var"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
	} `json:"links"`
}

type AppPermission struct {
	ReadBasicData     bool `json:"read_basic_data"`
	ReadSensitiveData bool `json:"read_sensitive_data"`
}

type AppSetDroplet struct {
	Data struct {
		GUID string `json:"guid" validate:"required"`
	} `json:"data" validate:"required"`
}

type AppSSH struct {
	Enabled bool   `json:"enabled"`
	Reason  string `json:"reason"`
}

type AppEnvVar struct {
	Var struct {
	} `json:"var" validate:"required"`
}
