package processes

import "time"

type Process struct {
	GUID        string `json:"guid"`
	Type        string `json:"type"`
	Command     string `json:"command"`
	Instances   int    `json:"instances"`
	MemoryInMb  int    `json:"memory_in_mb"`
	DiskInMb    int    `json:"disk_in_mb"`
	HealthCheck struct {
		Type string `json:"type"`
		Data struct {
			Timeout interface{} `json:"timeout"`
		} `json:"data"`
	} `json:"health_check"`
	Relationships struct {
		App struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"app"`
		Revision struct {
			Data struct {
				GUID string `json:"guid"`
			} `json:"data"`
		} `json:"revision"`
	} `json:"relationships"`
	Metadata struct {
		Labels struct {
		} `json:"labels"`
		Annotations struct {
		} `json:"annotations"`
	} `json:"metadata"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Scale struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"scale"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
		Space struct {
			Href string `json:"href"`
		} `json:"space"`
		Stats struct {
			Href string `json:"href"`
		} `json:"stats"`
	} `json:"links"`
}

type HealthCheck struct {
	Type string `json:"type"`
	Data struct {
		Timeout interface{} `json:"timeout"`
		/*InvocationTimeout interface{} `json:"invocation_timeout"`
		Endpoint string `json:"endpoint"`*/
	} `json:"data"`
}

type ProcessStats struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
	State string `json:"state"`
	Usage struct {
		Time time.Time `json:"time"`
		CPU  float64   `json:"cpu"`
		Mem  int       `json:"mem"`
		Disk int       `json:"disk"`
	} `json:"usage"`
	Host          string `json:"host"`
	InstancePorts []struct {
		External             int `json:"external"`
		Internal             int `json:"internal"`
		ExternalTLSProxyPort int `json:"external_tls_proxy_port"`
		InternalTLSProxyPort int `json:"internal_tls_proxy_port"`
	} `json:"instance_ports"`
	Uptime           int         `json:"uptime"`
	MemQuota         int         `json:"mem_quota"`
	DiskQuota        int         `json:"disk_quota"`
	FdsQuota         int         `json:"fds_quota"`
	IsolationSegment string      `json:"isolation_segment"`
	Details          interface{} `json:"details"`
}
