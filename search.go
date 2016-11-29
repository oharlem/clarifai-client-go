package clarifai

import (
	"net/url"
	"strconv"
)

const (
	DefaultItemsPerPageQty = 20
	SearchQueryTypeAnd     = "and"
	// SearchQueryTypeOr      = "or" // reserved
)

type Hit struct {
	Score float64 `json:"score"`
	Input *Input  `json:"input,omitempty"`
}

type SearchQuery struct {
	QueryObject *QueryObject `json:"query"`
	Session     *Session     `json:"-"`
	Page        int          `json:"-"`
	PerPage     int          `json:"-"`
	Type        string       `json:"-"`
}

func (q *SearchQuery) GetPage() int {
	return q.Page
}

func (q *SearchQuery) GetPerPage() int {
	return q.PerPage
}

// NewSearchQuery initializes a new query. Available types - "AND" and "OR".
func NewSearchQuery(t string) *SearchQuery {
	return &SearchQuery{
		Page:        0,
		PerPage:     DefaultItemsPerPageQty,
		QueryObject: &QueryObject{},
		Type:        t,
	}
}

// WithUserConcept adds a positive match condition to the user-defined set of concepts.
func (q *SearchQuery) WithUserConcept(s string) {

	i := Input{}
	i.AddConcept(s, true)

	qf := QueryFragment{
		Input: &i,
	}

	q.AddFragment(&qf)
}

// WithoutUserConcept adds a negative match condition to the user-defined set of concepts.
func (q *SearchQuery) WithoutUserConcept(s string) {

	i := Input{}
	i.AddConcept(s, false)

	qf := QueryFragment{
		Input: &i,
	}

	q.AddFragment(&qf)
}

// WithAPIConcept adds a positive match condition to the API-defined set of concepts.
func (q *SearchQuery) WithAPIConcept(s string) {

	qo := QueryOutput{}
	qo.AddConcept(s, true)

	qf := QueryFragment{
		Output: &qo,
	}

	q.AddFragment(&qf)
}

// WithoutAPIConcept adds a negative match condition to the API-provided set of concepts.
func (q *SearchQuery) WithoutAPIConcept(s string) {

	qo := QueryOutput{}
	qo.AddConcept(s, false)

	qf := QueryFragment{
		Output: &qo,
	}

	q.AddFragment(&qf)
}

// WithImage adds a positive match condition by an image.
func (q *SearchQuery) WithImage(i *Image) error {

	qf := QueryFragment{
		Output: &QueryOutput{
			Input: &Input{
				Data: i,
			},
		},
	}

	q.AddFragment(&qf)

	return nil
}

// WithMetadata adds a match filter for inputs, that were added with custom metadata.
func (q *SearchQuery) WithMetadata(m interface{}) {
	i := &Input{}
	i.SetMetadata(m)

	qf := QueryFragment{
		Input: i,
	}

	q.AddFragment(&qf)
}

// AddFragment adds fragment to the current clause of the query.
func (q *SearchQuery) AddFragment(qf *QueryFragment) {
	if q.Type == SearchQueryTypeAnd {
		q.QueryObject.QueryAnds = append(q.QueryObject.QueryAnds, qf)
	}
}

// SetPagination sets pagination configuration to the search query.
func (q *SearchQuery) SetPagination(page, perPage int) {
	q.Page = page
	q.PerPage = perPage
}

// Search issues a search request to Clarifai API with a provided search query.
func (s *Session) Search(q *SearchQuery) (*Response, error) {
	out := &Response{}

	payload, err := PrepPayload(q)
	if err != nil {
		return out, err
	}

	err = s.PostCall(GetEndpoint(s, "searches", q), payload, out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// GetEndpoint generates a service endpoint. First page starts with 1.
func GetEndpoint(s *Session, endpoint string, q Requester) string {

	uri := s.GetAPIHost(endpoint)

	if q.GetPage() > 0 && q.GetPerPage() > 0 {
		v := url.Values{}
		v.Set("page", strconv.Itoa(q.GetPage()))
		v.Add("per_page", strconv.Itoa(q.GetPerPage()))
		return uri + "?" + v.Encode()
	}

	return uri
}
