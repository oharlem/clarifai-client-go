package clarifai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	PredictInputLimit = 128
)

type PredictResponse struct {
	Status  *PredictStatus `json:"status,omitempty"`
	Outputs []*Output      `json:"outputs,omitempty"`
}

type Output struct {
	ID        string        `json:"id"`
	Status    PredictStatus `json:"status"`
	CreatedAt string        `json:"created_at"`
	Model     Model         `json:"model"`
	Input     Input         `json:"input"`
	Data      OutputData    `json:"data"`
}

type OutputData struct {
	Concepts *OutputConcepts `json:"concepts,omitempty"`
	Image    *ImageData      `json:"image,omitempty"`
	Metadata *Metadata       `json:"metadata,omitempty"`
}

type OutputConcepts []ResponseConcept

type PredictStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type PredictService struct {
	Inputs  InputsCollection
	ModelID string
	Session *Session
}

func NewPredictService(s *Session) *PredictService {
	return &PredictService{
		ModelID: PublicModelGeneral,
		Session: s,
	}
}

// AddInput adds input to the predict request.
// Limited to 128 inputs per call as per https://developer-preview.clarifai.com/guide/inputs#inputs.
func (svc *PredictService) AddInput(i *Input) error {
	if svc.GetInputsQty() >= PredictInputLimit {
		return ErrInputLimitReached
	}
	svc.Inputs.Inputs = append(svc.Inputs.Inputs, i)
	return nil
}

func (svc *PredictService) GetInputsQty() int {
	return len(svc.Inputs.Inputs)
}

// SetModel is an optional model setter for predict calls. Defaults to the general model.
func (svc *PredictService) SetModel(modelID string) {
	svc.ModelID = modelID
}

func (svc *PredictService) GetModel() string {
	return svc.ModelID
}

// GetModelEndpoint generates a model endpoint based on currently set model ID.
func (svc *PredictService) GetModelEndpoint() string {
	return getAPIHost("models/" + svc.GetModel() + "/outputs")
}

// Call sends a predicate request to the API.
func (svc *PredictService) Call() (*PredictResponse, error) {

	out := &PredictResponse{}

	err := svc.CallValidations()
	if err != nil {
		return out, err
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(svc.Inputs)

	req, err := http.NewRequest(http.MethodPost, svc.GetModelEndpoint(), b)
	if err != nil {
		return out, err
	}
	req.Header.Set("Authorization", "Bearer "+svc.Session.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return out, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return out, err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (svc *PredictService) CallValidations() error {

	if svc.GetInputsQty() == 0 {
		return ErrNoInputs
	}

	return nil
}
