package tasks

import "time"

//*struct ,omitempty" post로 보낼 객체들만 선언

type Tasks struct {
	GUID       string `json:"guid"`
	SequenceID int    `json:"sequence_id"`
	Name       string `json:"name"`
	Command    string `json:"command"`
	State      string `json:"state"`
	MemoryInMb int    `json:"memory_in_mb"`
	DiskInMb   int    `json:"disk_in_mb"`
	Result     struct {
		FailureReason interface{} `json:"failure_reason"`
	} `json:"result"`
	DropletGUID string `json:"droplet_guid"`
	Metadata    struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
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
		App struct {
			Href string `json:"href"`
		} `json:"app"`
		Cancel struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"cancel"`
		Droplet struct {
			Href string `json:"href"`
		} `json:"droplet"`
	} `json:"links"`
}

type CreateTask struct {
	Name        string `json:"name,omitempty"`
	Command     string `json:"command,omitempty"`
	DropletGUID string `json:"droplet_guid,omitempty"`
	Metadata    *struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata,omitempty"`
	Template *struct {
		Process struct {
			GUID string `json:"guid"`
		} `json:"process"`
	} `json:"template,omitempty"`
}

type AppTask struct {
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
	Resources []Tasks `json:"resources"`
}

type UpdateTask struct {
	Metadata struct {
		Labels struct {
			Key string `json:"key"`
		} `json:"labels"`
		Annotations struct {
			Note string `json:"note"`
		} `json:"annotations"`
	} `json:"metadata"`
}
