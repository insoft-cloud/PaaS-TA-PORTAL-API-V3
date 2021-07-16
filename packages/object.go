package packages

import "time"

type Package struct {
	GUID string `json:"guid"`
	Type string `json:"type"`
	Data struct {
		Checksum struct {
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"checksum"`
		Error interface{} `json:"error"`
	} `json:"data"`
	State         string    `json:"state"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
		Upload struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"upload"`
		Download struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"download"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
	} `json:"links"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
}

type CreatePackage struct {
	Type          string `json:"type" validate:"required"`
	Relationships *struct {
		App struct {
			Data struct {
				GUID []string `json:"guid"`
			} `json:"data"`
		} `json:"app" validate:"required"`
	} `json:"relationships"`
	Data *struct {
		Checksum struct {
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"checksum"`
		Image    string `json:"image"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"data,omitempty"`
	Metadata *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type PackageList struct {
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
	Resources []Package `json:"resources"`
}

type UpdatePackage struct {
	Metadata struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
}

type CopyPackage struct {
	Relationships struct {
		App struct {
			Data struct {
				GUID []string `json:"guid"`
			} `json:"data"`
		} `json:"app"`
	} `json:"relationships" validate:"required"`
}

type Resources struct {
	Resources []struct {
		Checksum struct {
			Value string `json:"value"`
		} `json:"checksum"`
		SizeInBytes int    `json:"size_in_bytes"`
		Path        string `json:"path"`
		Mode        string `json:"mode"`
	} `json:"resources"`
}
