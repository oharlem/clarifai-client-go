package clarifai

// Output query fragment.
type Output struct {
	ID        string         `json:"id"`
	Status    *ServiceStatus `json:"status,omitempty"`
	CreatedAt string         `json:"created_at"`
	Model     *Model         `json:"model,omitempty"`
	Input     *Input         `json:"input,omitempty"`
	Data      *OutputData    `json:"data,omitempty"`
}

type OutputData struct {
	Concepts *OutputConcepts `json:"concepts,omitempty"`
	Image    *ImageData      `json:"image,omitempty"`
	Metadata *interface{}    `json:"metadata,omitempty"`
}

type OutputConcepts []*OutputConcept

type OutputConcept struct {
	AppID *string `json:"app_id,omitempty"`
	ID    string  `json:"id"`
	Name  string  `json:"name,omitempty"`
	Value float64 `json:"value,omitempty"`
}
