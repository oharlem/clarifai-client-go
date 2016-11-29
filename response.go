package clarifai

type Response struct {
	Status  *ServiceStatus `json:"status,omitempty"`
	Outputs []*Output      `json:"outputs,omitempty"`
	Input   *Input         `json:"input,omitempty"` // Request for one input.
	Inputs  []*Input       `json:"inputs,omitempty"`
	Hits    []*Hit         `json:"hits,omitempty"` // Search hits.
}
