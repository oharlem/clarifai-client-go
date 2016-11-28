package clarifai

// Requester is an interface for all request-type objects: Request, SearchQuery.
type Requester interface {
	GetPage() int
	GetPerPage() int
}

type Request struct {
	Inputs  []*Input `json:"inputs"`
	ModelID string   `json:"-"`
	Page    int      `json:"-"`
	PerPage int      `json:"-"`
}

func (r *Request) GetPage() int {
	return r.Page
}

func (r *Request) GetPerPage() int {
	return r.PerPage
}

// WithPagination sets pagination configuration to the request.
func (r *Request) WithPagination(page, perPage int) {
	r.Page = page
	r.PerPage = perPage
}

func (r *Request) GetModel() string {
	return r.ModelID
}

func (r *Request) AddInput(i *Input) error {
	if len(r.Inputs) >= InputLimit {
		return ErrInputLimitReached
	}
	r.Inputs = append(r.Inputs, i)
	return nil
}

// AddImageInput adds an image input to a request.
func (r *Request) AddImageInput(i *Image) error {
	if len(r.Inputs) >= InputLimit {
		return ErrInputLimitReached
	}

	in := &Input{
		Data: i,
	}

	r.Inputs = append(r.Inputs, in)
	return nil
}

// SetModel is an optional model setter for predict calls. Defaults to the general model.
func (r *Request) SetModel(modelID string) {
	r.ModelID = modelID
}

func NewRequest(s *Session) *Request {
	return &Request{
		ModelID: PublicModelGeneral,
		Page:    0,
		PerPage: DefaultItemsPerPageQty,
	}
}
