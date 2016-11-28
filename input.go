package clarifai

var AllowedMimeTypes map[string]struct{}

func init() {
	AllowedMimeTypes = map[string]struct{}{
		"image/bmp":  struct{}{},
		"image/jpeg": struct{}{},
		"image/png":  struct{}{},
		"image/tiff": struct{}{},
	}
}

type Input struct {
	Data      *Image         `json:"data,omitempty"`
	ID        string         `json:"id,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	Status    *ServiceStatus `json:"status,omitempty"`
}

type InputObject struct {
	Inputs []*Input `json:"inputs"`
}

func (o *InputObject) AddInput(i *Input) {
	o.Inputs = append(o.Inputs, i)
}

func (o *InputObject) GetInputsQty() int {
	if o.Inputs == nil {
		return 0
	}
	return len(o.Inputs)
}

//
//
// API
//
//

// AddImagesToIndex sends images to a search index.
func (s *Session) AddImagesToIndex(r *Request) (*Response, error) {
	out := &Response{}

	payload, err := PrepPayload(r)
	if err != nil {
		return out, err
	}

	err = s.PostCall(GetEndpoint(s, "inputs", r), payload, out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// Get a list of all inputs.
func (s *Session) ListAllInputs(r *Request) (*Response, error) {
	out := &Response{}

	err := s.GetCall(GetEndpoint(s, "inputs", r), out)
	if err != nil {
		return out, err
	}

	return out, nil
}
