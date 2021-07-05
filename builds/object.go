package builds

import "time"

type Build struct {
	GUID      string    `json:"guid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy struct {
		GUID  string `json:"guid"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"created_by"`
	State     string      `json:"state"`
	Error     interface{} `json:"error"`
	Lifecycle struct {
		Type string `json:"type"`
		Data struct {
			Buildpacks []string `json:"buildpacks"`
			Stack      string   `json:"stack"`
		} `json:"data"`
	} `json:"lifecycle"`
	Package struct {
		GUID string `json:"guid"`
	} `json:"package"`
	Droplet struct {
		GUID string `json:"guid"`
	} `json:"droplet"`
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app"`
	} `json:"relationships"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
		Droplet struct {
			Href string `json:"href"`
		} `json:"droplet"`
	} `json:"links"`
}

type BuildList struct {
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
	Resources []Build `json:"resources"`
}

type CreateBuild struct {
	Package struct {
		GUID string `json:"guid"`
	} `json:"package" validate:"required"`
	Lifecycle struct {
		Type string `json:"type"`
		Data struct {
			Buildpacks []string `json:"buildpacks"`
			Stack      string   `json:"stack"`
		} `json:"data"`
	} `json:"lifecycle"`
	StagingMemoryInMb int `json:"staging_memory_in_mb"`
	StagingDiskInMb   int `json:"staging_disk_in_mb"`
	Metadata          struct {
		Labels struct {
			Environment    string `json:"environment"`
			InternetFacing string `json:"internet-facing"`
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
}

type UpdateBuild struct {
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
		Lifecycle struct {
			Data struct {
				image string `json:"image"`
			} `json:"data"`
		} `json:"lifecycle"`
		State string `json:"state"`
	} `json:"metadata"`
}
