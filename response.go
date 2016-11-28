package clarifai

type Response struct {
	Status  *ServiceStatus `json:"status,omitempty"`
	Outputs []*Output      `json:"outputs,omitempty"`
	Inputs  []*Input       `json:"inputs,omitempty"`
	Hits    []*Hit         `json:"hits,omitempty"` // search hits
}
