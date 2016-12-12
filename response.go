package clarifai

// Response is a universal Clarifai API response object.
type Response struct {
	Status        *ServiceStatus  `json:"status,omitempty"`
	Outputs       []*Output       `json:"outputs,omitempty"`
	Input         *Input          `json:"input,omitempty"` // Request for one input.
	Inputs        []*Input        `json:"inputs,omitempty"`
	Hits          []*Hit          `json:"hits,omitempty"` // Search hits.
	Model         *Model          `json:"model,omitempty"`
	Models        []*Model        `json:"models,omitempty"`
	ModelVersion  *ModelVersion   `json:"model_version,omitempty"`
	ModelVersions []*ModelVersion `json:"model_versions,omitempty"`
}
