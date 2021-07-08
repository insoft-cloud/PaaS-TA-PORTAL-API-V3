package resource_matches

type ResourceMatch struct {
	Resources []struct {
		Checksum struct {
			Value string `json:"value" validate:"required"`
		} `json:"checksum"`
		SizeInBytes int    `json:"size_in_bytes" validate:"required"`
		Path        string `json:"path"`
		Mode        string `json:"mode"`
	} `json:"resources"`
}
