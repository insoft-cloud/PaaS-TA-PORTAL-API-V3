package domains

import (
	"time"
)

type CreateDomain struct {
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

type Domain struct {
	GUID               string      `json:"guid"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	Name               string      `json:"name"`
	Internal           bool        `json:"internal"`
	RouterGroup        interface{} `json:"router_group"`
	SupportedProtocols []string    `json:"supported_protocols"`
	Metadata           struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Relationships struct {
		Organization struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organization"`
		SharedOrganizations struct {
			Data []struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"shared_organizations"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
		RouteReservations struct {
			Href string `json:"href"`
		} `json:"route_reservations"`
		SharedOrganizations struct {
			Href string `json:"href"`
		} `json:"shared_organizations"`
	} `json:"links"`
}

type DomainList struct {
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
	Resources []Domain `json:"resources"`
}

type OrganizationDomainsList struct {
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
		GUID        string    `json:"guid"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Name        string    `json:"name"`
		Internal    bool      `json:"internal"`
		RouterGroup struct {
			GUID string `json:"guid"`
		} `json:"router_group"`
		SupportedProtocols []string `json:"supported_protocols"`
		Metadata           struct {
			Labels struct {
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
		} `json:"metadata"`
		Relationships struct {
			Organization struct {
				Data interface{} `json:"data"`
			} `json:"organization"`
			SharedOrganizations struct {
				Data []interface{} `json:"data"`
			} `json:"shared_organizations"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			RouteReservations struct {
				Href string `json:"href"`
			} `json:"route_reservations"`
			RouterGroup struct {
				Href string `json:"href"`
			} `json:"router_group"`
		} `json:"links"`
	} `json:"resources"`
}

type UpdateDomains struct {
	GUID               string      `json:"guid"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	Name               string      `json:"name"`
	Internal           bool        `json:"internal"`
	RouterGroup        interface{} `json:"router_group"`
	SupportedProtocols []string    `json:"supported_protocols"`
	Metadata           struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata"`
	Relationships struct {
		Organization struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"organization"`
		SharedOrganizations struct {
			Data []struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"shared_organizations"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
		RouteReservations struct {
			Href string `json:"href"`
		} `json:"route_reservations"`
		SharedOrganizations struct {
			Href string `json:"href"`
		} `json:"shared_organizations"`
	} `json:"links"`
}

type ShareDomains struct {
	Data []struct {
		GUID string `json:"guid" validate:"required"`
	} `json:"data" `
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
