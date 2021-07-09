package routes

import "time"

type Route struct {
	GUID         string    `json:"guid"`
	Protocol     string    `json:"protocol"`
	Port         int       `json:"port"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Host         string    `json:"host"`
	Path         string    `json:"path"`
	URL          string    `json:"url"`
	Destinations []struct {
		GUID string `json:"guid"`
		App  struct {
			GUID    string `json:"guid"`
			Process struct {
				Type string `json:"type"`
			} `json:"process"`
		} `json:"app"`
		Weight interface{} `json:"weight"`
		Port   int         `json:"port"`
	} `json:"destinations"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Relationships struct {
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
		Domain struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"domain"`
	} `json:"relationships"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Space struct {
			Href string `json:"href"`
		} `json:"space"`
		Domain struct {
			Href string `json:"href"`
		} `json:"domain"`
		Destinations struct {
			Href string `json:"href"`
		} `json:"destinations"`
	} `json:"links"`
}

type Destination struct {
	GUID string `json:"guid"`
	App  struct {
		GUID    string `json:"guid"`
		Process struct {
			Type string `json:"type"`
		} `json:"process"`
	} `json:"app"`
	Weight interface{} `json:"weight"`
	Port   int         `json:"port"`
}

type CreateRoute struct {
	Host          string `json:"host,omitempty"`
	Path          string `json:"path,omitempty"`
	Port          int    `json:"port,omitempty"`
	Relationships *struct {
		Domain struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"domain"`
		Space struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"space"`
	} `json:"relationships" validate:"required"`
	Metadata *struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type RouteList struct {
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
	Resources []Route `json:"resources"`
}

type UpdateRoute struct {
	Metadata struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type insertDestinations struct {
	Destinations []struct {
		App struct {
			GUID    string `json:"guid"`
			Process struct {
				Type string `json:"type"`
			} `json:"process"`
		} `json:"app,omitempty"`
		Port int `json:"port,omitempty"`
	} `json:"destinations"`
}

type ReplaceAllDestinationRoute struct {
	Destinations []struct {
		Weight int `json:"weight"`
		App    struct {
			GUID    string `json:"guid"`
			Process struct {
				Type string `json:"type"`
			} `json:"process"`
		} `json:"app,omitempty"`
		Port int `json:"port,omitempty"`
	} `json:"destinations" validate:"required"`
}
