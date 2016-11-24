package clarifai

type PredictResponse struct {
	Status  *ServiceStatus `json:"status,omitempty"`
	Outputs []*Output      `json:"outputs,omitempty"`
}

func (r *PredictResponse) GetStatus() *ServiceStatus {
	return r.Status
}

type PredictService struct {
	InputObject *InputObject
	ModelID     string
	Session     *Session
	Page        int
	PerPage     int
}

func (svc *PredictService) GetPage() int {
	return svc.Page
}

func (svc *PredictService) GetPerPage() int {
	return svc.PerPage
}

func (svc *PredictService) GetSession() *Session {
	return svc.Session
}

func (svc *PredictService) GetInputObject() *InputObject {
	return svc.InputObject
}

// NewPredictService returns a default predict service object.
func NewPredictService(s *Session) *PredictService {
	return &PredictService{
		ModelID:     PublicModelGeneral,
		Session:     s,
		InputObject: &InputObject{},
	}
}

// AddInput adds input to the predict request.
// Limited to 128 inputs per call as per https://developer-preview.clarifai.com/guide/inputs#inputs.
func (svc *PredictService) AddInput(i *Input) error {
	if svc.InputObject.GetInputsQty() >= InputLimit {
		return ErrInputLimitReached
	}
	svc.InputObject.AddInput(i)
	return nil
}

// SetModel is an optional model setter for predict calls. Defaults to the general model.
func (svc *PredictService) SetModel(modelID string) {
	svc.ModelID = modelID
}

// GetModel return currently assigned model of a service.
func (svc *PredictService) GetModel() string {
	return svc.ModelID
}

// GetPredictions sends a predict request to the API.
func (svc *PredictService) GetPredictions() (*PredictResponse, error) {
	out := &PredictResponse{}

	// Inputs for this call are required.
	if svc.InputObject.GetInputsQty() == 0 {
		return out, ErrNoInputs
	}

	// Generates a request-specific endpoint based on currently set model ID.
	endpoint := GetURI(svc, "models/"+svc.GetModel()+"/outputs")

	payload, err := GetPayload(svc)
	if err != nil {
		return out, err
	}

	err = PostCall(svc, endpoint, payload, out)
	if err != nil {
		return out, err
	}

	return out, nil

}
