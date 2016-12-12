package clarifai

import (
	"net/http"
)

const (
	// Public pre-defined models.
	// Source: https://developer-preview.clarifai.com/guide/publicmodels#public-models

	// PublicModelGeneral is a public model "general".
	PublicModelGeneral = "aaa03c23b3724a16a56b629203edc62c"

	// PublicModelFood is a public model "food".
	PublicModelFood = "bd367be194cf45149e75f01d59f77ba7"

	// PublicModelTravel is a public model "travel".
	PublicModelTravel = "eee28c313d69466f836ab83287a54ed9"

	// PublicModelNSFW is a public model "NSFW" (Not Safe For Work).
	PublicModelNSFW = "e9576d86d2004ed1a38ba0cf39ecb4b1"

	// PublicModelWeddings is a public model "weddings".
	PublicModelWeddings = "c386b7a870114f4a87477c0824499348"

	// PublicModelColor is a public model "color".
	PublicModelColor = "eeed0b6733a644cea07cf4c60f87ebb7"
)

type ModelRequest struct {
	Model Model `json:"model"`
}

type Model struct {
	Name         *string       `json:"name,omitempty"`
	ID           *string       `json:"id,omitempty"`
	CreatedAt    *string       `json:"created_at,omitempty"`
	AppID        *string       `json:"app_id,omitempty"`
	OutputInfo   *OutputInfo   `json:"output_info,omitempty"`
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

type OutputInfo struct {
	Message      string        `json:"message,omitempty"`
	Type         string        `json:"type,omitempty"`
	OutputConfig *OutputConfig `json:"output_config,omitempty"`
	OutputData   *OutputData   `json:"data,omitempty"`
}

type ModelVersion struct {
	ID        string         `json:"id"`
	CreatedAt string         `json:"created_at"`
	Status    *ServiceStatus `json:"status"`
}

type OutputConfig struct {
	ConceptsMutuallyExclusive bool `json:"concepts_mutually_exclusive"`
	ClosedEnvironment         bool `json:"closed_environment"`
}

// modelOptions is a model configuration object used to set optional settings for a new model.
type modelOptions struct {
	ID                        string   // Model ID. If not set, wil be generated automatically.
	Concepts                  []string // Optional concepts to associated with this model
	ConceptsMutuallyExclusive bool     // True or False, whether concepts are mutually exclusive
	ClosedEnvironment         bool     // True or False, whether use negatives for prediction
}

// ModelQuery is a model search payload.
type ModelQuery struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// NewModelOptions returns default model options.
func NewModelOptions() *modelOptions {
	return &modelOptions{}
}

// Predict fetches prediction info for a provided asset from a given model.
func (s *Session) Predict(i *Inputs) *Request {

	r := NewRequest(s, http.MethodPost, "models/"+i.modelID+"/outputs")
	r.SetPayload(i)

	return r
}

// CreateModel creates a new model. If ID is empty, it will be created automatically by Clarifai API.
func (s *Session) CreateModel(name string, opt *modelOptions) *Request {

	p := ModelRequest{
		Model: Model{
			Name: &name,
		},
	}

	if p.Model.OutputInfo == nil {
		p.Model.OutputInfo = &OutputInfo{
			OutputData:   &OutputData{},
			OutputConfig: &OutputConfig{},
		}
	}

	if opt.ID != "" {
		p.Model.ID = &opt.ID
	}

	if len(opt.Concepts) > 0 {
		var concepts []*OutputConcept

		for _, id := range opt.Concepts {
			concepts = append(concepts, &OutputConcept{
				ID: id,
			})
		}

		p.Model.OutputInfo.OutputData.Concepts = concepts
	}

	if opt.Concepts != nil {
		p.Model.OutputInfo.OutputConfig.ConceptsMutuallyExclusive = opt.ConceptsMutuallyExclusive
	}

	if opt.Concepts != nil {
		p.Model.OutputInfo.OutputConfig.ClosedEnvironment = opt.ClosedEnvironment
	}

	r := NewRequest(s, http.MethodPost, "models")
	r.SetPayload(p)

	return r
}

// AddModelConcepts adds new concepts to an existing model.
func (s *Session) AddModelConcepts(ID string, c []string) *Request {

	r := NewRequest(s, http.MethodPatch, "models/"+ID+"/output_info/data/concepts")

	p := struct {
		Concepts []*OutputConcept `json:"concepts"`
		Action   string           `json:"action"`
	}{
		Concepts: sliceToConcepts(c),
		Action:   "merge_concepts",
	}

	r.SetPayload(p)

	return r
}

// GetModels fetches a list of all models, including custom and public.
func (s *Session) GetModels() *Request {

	return NewRequest(s, http.MethodGet, "models")
}

// GetModel fetches a single model by its ID.
func (s *Session) GetModel(ID string) *Request {

	return NewRequest(s, http.MethodGet, "models/"+ID)
}

// GetModelOutput fetches a single model by its ID with output_info data.
func (s *Session) GetModelOutput(ID string) *Request {

	return NewRequest(s, http.MethodGet, "models/"+ID+"/output_info")
}

// GetModelVersion fetches version data of a single model .
func (s *Session) GetModelVersion(m, v string) *Request {

	return NewRequest(s, http.MethodGet, "models/"+m+"/versions/"+v)
}

// GetModelVersionInputs fetches inputs used to train a specific model version.
func (s *Session) GetModelVersionInputs(m, v string) *Request {

	return NewRequest(s, http.MethodGet, "models/"+m+"/versions/"+v+"/inputs")
}

// GetModelVersions fetches a single model by its ID with versions data.
func (s *Session) GetModelVersions(ID string) *Request {

	return NewRequest(s, http.MethodGet, "models/"+ID+"/versions")
}

// GetModelInputs fetches inputs of a single model by its ID.
func (s *Session) GetModelInputs(ID string) *Request {

	return NewRequest(s, http.MethodGet, "models/"+ID+"/inputs")
}

// DeleteModelVersion deletes a specific version of a model.
func (s *Session) DeleteModelVersion(m, v string) *Request {

	return NewRequest(s, http.MethodDelete, "models/"+m+"/versions/"+v)
}

// DeleteModel deletes a single model by ID.
func (s *Session) DeleteModel(ID string) *Request {

	return NewRequest(s, http.MethodDelete, "models/"+ID)
}

// DeleteAllModels deletes all models associated with your application.
func (s *Session) DeleteAllModels() *Request {

	return NewRequest(s, http.MethodDelete, "models")
}

// TrainModel starts a model training operation.
// When you train a model, you are telling the system to look at all the images with concepts you've provided and learn from them.
// This train operation is asynchronous. It may take a few seconds for your model to be fully trained and ready.
func (s *Session) TrainModel(ID string) *Request {

	r := NewRequest(s, http.MethodPost, "models/"+ID+"/versions")

	return r
}

// DeleteModelConcepts removes concepts from a model.
func (s *Session) DeleteModelConcepts(ID string, c []string) *Request {

	r := NewRequest(s, http.MethodPatch, "models/"+ID+"/output_info/data/concepts")

	p := struct {
		Concepts []*OutputConcept `json:"concepts"`
		Action   string           `json:"action"`
	}{
		Concepts: sliceToConcepts(c),
		Action:   "delete_concepts",
	}

	r.SetPayload(p)

	return r
}

// SearchModel searches models by name and/or type.
func (s *Session) SearchModel(n, t string) *Request {

	r := NewRequest(s, http.MethodPost, "models/searches")

	mq := ModelQuery{}
	if n != "" {
		mq.Name = n
	}
	if t != "" {
		mq.Type = t
	}

	p := struct {
		ModelQuery ModelQuery `json:"model_query"`
	}{
		ModelQuery: mq,
	}

	r.SetPayload(p)

	return r
}

// Convert a slice of concepts into a request format.
func sliceToConcepts(c []string) []*OutputConcept {
	var concepts []*OutputConcept

	for _, id := range c {
		concepts = append(concepts, &OutputConcept{
			ID: id,
		})
	}

	return concepts
}
