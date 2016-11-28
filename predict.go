package clarifai

func (s *Session) GetPredictions(r *Request) (*Response, error) {

	out := &Response{}

	payload, err := PrepPayload(r)
	if err != nil {
		return out, err
	}

	err = s.PostCall(GetEndpoint(s, "models/"+r.GetModel()+"/outputs", r), payload, out)
	if err != nil {
		return out, err
	}

	return out, nil
}