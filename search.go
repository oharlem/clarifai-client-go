package clarifai

import "net/http"

const (
	SearchQueryTypeAnd = "and"
	// SearchQueryTypeOr      = "or" // reserved
)

type Hit struct {
	Score float64 `json:"score"`
	Input *Input  `json:"input,omitempty"`
}

type SearchRequest struct {
	QueryObject *QueryObject `json:"query,omitempty"`
	Type        string       `json:"-"`
	Pagination  *pagination  `json:"pagination,omitempty"`
}

type pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// NewSearchQuery initializes a new SearchRequest object with search query properties.
func NewSearchQuery(t string) *SearchRequest {
	return &SearchRequest{
		QueryObject: &QueryObject{},
		Type:        t,
	}
}

// NewAndSearchQuery for a search query of type "and"
func NewAndSearchQuery() *SearchRequest {
	return NewSearchQuery(SearchQueryTypeAnd)
}

// WithUserConcept adds a positive match condition to the user-defined set of concepts.
func (r *SearchRequest) WithUserConcept(c string) {

	i := Input{}
	i.AddConcept(c, true)

	qf := QueryFragment{
		Input: &i,
	}

	r.addFragment(&qf)
}

// WithoutUserConcept adds a negative match condition to the user-defined set of concepts.
func (r *SearchRequest) WithoutUserConcept(c string) {

	i := Input{}
	i.AddConcept(c, false)

	qf := QueryFragment{
		Input: &i,
	}

	r.addFragment(&qf)
}

// WithAPIConcept adds a positive match condition to the API-defined set of concepts.
func (r *SearchRequest) WithAPIConcept(c string) {

	qo := QueryOutput{}
	qo.AddConcept(c, true)

	qf := QueryFragment{
		Output: &qo,
	}

	r.addFragment(&qf)
}

// WithoutAPIConcept adds a negative match condition to the API-provided set of concepts.
func (r *SearchRequest) WithoutAPIConcept(c string) {

	qo := QueryOutput{}
	qo.AddConcept(c, false)

	qf := QueryFragment{
		Output: &qo,
	}

	r.addFragment(&qf)
}

// WithImage adds a positive match condition by an image.
func (r *SearchRequest) WithImage(i *Image) {

	qf := QueryFragment{
		Output: &QueryOutput{
			Input: &Input{
				Data: i,
			},
		},
	}

	r.addFragment(&qf)
}

// WithMetadata adds a match filter for inputs, that were added with custom metadata.
func (r *SearchRequest) WithMetadata(m interface{}) {
	i := &Input{}
	i.SetMetadata(m)

	qf := QueryFragment{
		Input: i,
	}

	r.addFragment(&qf)
}

// addFragment adds fragment to the current clause of the query.
func (r *SearchRequest) addFragment(qf *QueryFragment) {
	if r.Type == SearchQueryTypeAnd {
		r.QueryObject.Ands = append(r.QueryObject.Ands, qf)
	}
}

// Search issues a search request to Clarifai API with a provided search query.
func (s *Session) Search(p *SearchRequest) *Request {

	r := NewRequest(s, http.MethodPost, "searches")
	r.SetPayload(p)

	return r
}
