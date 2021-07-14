package manifests

type Diff struct {
	Diff []struct {
		Op   string `json:"op"`
		Path string `json:"path"`
		Was  struct {
			Route string `json:"route"`
		} `json:"was,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"diff"`
}
