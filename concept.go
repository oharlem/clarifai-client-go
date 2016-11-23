package clarifai

type InputConcept struct {
	ID    string `json:"id"`
	Value bool   `json:"value,omitempty"`
}

type ResponseConcept struct {
	AppID string  `json:"app_id,omitempty"`
	ID    string  `json:"id"`
	Name  string  `json:"name,omitempty"`
	Value float64 `json:"value,omitempty"`
}
