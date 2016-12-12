package clarifai

import "net/http"

type Input struct {
	Data      *Image         `json:"data,omitempty"`
	ID        string         `json:"id,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	Status    *ServiceStatus `json:"status,omitempty"`
}

type Inputs struct {
	Inputs  []*Input `json:"inputs"`
	modelID string   `json:"-"`
}

// InitInputs returns a default inputs object.
func InitInputs() *Inputs {
	return &Inputs{
		modelID: PublicModelGeneral,
	}
}

// AddInput adds an image input to a request.
func (i *Inputs) AddInput(im *Image, id string) error {
	if len(i.Inputs) >= InputLimit {
		return ErrInputLimitReached
	}

	in := &Input{
		Data: im,
	}

	// Add custom ID if provided.
	if id != "" {
		in.ID = id
	}

	i.Inputs = append(i.Inputs, in)
	return nil
}

// SetModel is an optional model setter for predict calls.
func (i *Inputs) SetModel(m string) {
	i.modelID = m
}

// AddConcept adds concepts to input.
func (i *Input) AddConcept(id string, value interface{}) {

	if i.Data == nil {
		i.Data = &Image{}
	}

	i.Data.Concepts = append(i.Data.Concepts, map[string]interface{}{
		"name":  id,
		"value": value,
	})
}

// SetMetadata adds metadata to a query input item ("input" -> "data" -> "metadata").
func (q *Input) SetMetadata(i interface{}) {
	if q.Data == nil {
		q.Data = &Image{}
	}
	q.Data.Metadata = i
}

// AddInputs builds a request to add inputs to the API.
func (s *Session) AddInputs(p *Inputs) *Request {

	r := NewRequest(s, http.MethodPost, "inputs")
	r.SetPayload(p)

	return r
}

// GetAllInputs fetches a list of all inputs.
func (s *Session) GetAllInputs() *Request {

	return NewRequest(s, http.MethodGet, "inputs")
}

// GetInput fetches one input.
func (s *Session) GetInput(id string) *Request {

	return NewRequest(s, http.MethodGet, "inputs/"+id)
}

// GetInputStatuses fetches statuses of all inputs.
func (s *Session) GetInputStatuses() *Request {

	return NewRequest(s, http.MethodGet, "inputs/status")
}

// DeleteInputConcepts remove concepts that were already added to an input.
func (s *Session) DeleteInputConcepts(id string, concepts []string) *Request {

	// 1. Build a request.
	r := NewRequest(s, http.MethodPatch, "inputs/"+id+"/data/concepts")

	// 2. Add payload.
	p := struct {
		Concepts []*OutputConcept `json:"concepts,omitempty"`
		Action   string           `json:"action"`
	}{}

	for _, v := range concepts {
		oc := &OutputConcept{
			ID: v,
		}
		p.Concepts = append(p.Concepts, oc)
	}
	p.Action = "delete_concepts"

	r.SetPayload(p)

	return r
}

// UpdateInputConcepts updates existing and/or adds new concepts to an input by its ID.
func (s *Session) UpdateInputConcepts(id string, userConcepts map[string]bool) *Request {

	// 1. Build a request.
	r := NewRequest(s, http.MethodPatch, "inputs/"+id+"/data/concepts")

	// 2. Add payload.
	// Convert an input map into a map of concepts.
	var reqConcepts []map[string]interface{}

	for id, value := range userConcepts {

		reqConcepts = append(reqConcepts, map[string]interface{}{
			"id":    id,
			"value": value,
		})
	}

	p := struct {
		Concepts []map[string]interface{} `json:"concepts"`
		Action   string                   `json:"action"`
	}{
		Concepts: reqConcepts,
		Action:   "merge_concepts",
	}

	r.SetPayload(p)

	return r
}

// DeleteInput deletes a single input by its ID.
func (s *Session) DeleteInput(id string) *Request {

	return NewRequest(s, http.MethodDelete, "inputs/"+id)
}

// DeleteInputs deletes multiple inputs by their IDs.
func (s *Session) DeleteInputs(ids []string) *Request {

	// 1. Build a request.
	r := NewRequest(s, http.MethodDelete, "inputs")

	// 2. Add a payload.
	r.SetPayload(struct {
		Inputs []string `json:"ids"`
	}{
		Inputs: ids,
	})

	return r
}

// DeleteAllInputs deletes all inputs.
func (s *Session) DeleteAllInputs() *Request {

	return NewRequest(s, http.MethodDelete, "inputs")
}
