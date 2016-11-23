package clarifai

const (
	// Public pre-defined models.
	// Source: https://developer-preview.clarifai.com/guide/publicmodels#public-models
	PublicModelGeneral  = "aaa03c23b3724a16a56b629203edc62c"
	PublicModelFood     = "bd367be194cf45149e75f01d59f77ba7"
	PublicModelTravel   = "eee28c313d69466f836ab83287a54ed9"
	PublicModelNSFW     = "e9576d86d2004ed1a38ba0cf39ecb4b1"
	PublicModelWeddings = "c386b7a870114f4a87477c0824499348"
	PublicModelColor    = "eeed0b6733a644cea07cf4c60f87ebb7"
)

type Model struct {
	Name         string       `json:"name"`
	ID           string       `json:"id"`
	CreatedAt    string       `json:"created_at"`
	AppID        string       `json:"app_id"`
	OutputInfo   OutputInfo   `json:"output_info"`
	ModelVersion ModelVersion `json:"model_version"`
}

type OutputInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type ModelVersion struct {
	ID        string      `json:"id"`
	CreatedAt string      `json:"created_at"`
	Status    ModelStatus `json:"status"`
}

type ModelStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
