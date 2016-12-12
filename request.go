package clarifai

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultPage            = 1
	defaultItemsPerPageQty = 20
)

// Request contains all information necessary to create an HTTP request to Clarifai API.
type Request struct {
	method  string
	page    int
	perPage int
	path    string
	payload interface{}
	session *Session
}

// NewRequest generates a new Request object with default settings.
func NewRequest(s *Session, method, path string) *Request {
	return &Request{
		method: method,
		// First page number is 1, not 0, but I use default 0
		// is an indicator of unset value for pagination logic.
		path:    path,
		page:    defaultPage,
		perPage: defaultItemsPerPageQty,
		session: s,
	}
}

// SetPayload sets payload to request.
func (r *Request) SetPayload(p interface{}) {
	r.payload = p
}

// WithPagination adds pagination configuration to the request.
func (r *Request) WithPagination(page, perPage int) *Request {
	r.page = page
	r.perPage = perPage

	return r
}

// Do sends a request to API.
func (r *Request) Do() (*Response, error) {

	var resp *Response
	var err error

	switch r.method {
	case http.MethodGet:
		r.addPagination()
		resp, err = r.session.HTTPCall(r.method, r.path, nil)
	case http.MethodPost, http.MethodPatch, http.MethodDelete:
		r.addPagination()
		resp, err = r.session.HTTPCall(r.method, r.path, r.payload)
	default:
		panic("Unsupported HTTP method!")
	}

	return resp, err
}

// addPagination adds pagination arguments to endpoint path.
func (r *Request) addPagination() {

	if r.page > 0 && r.perPage > 0 {

		// Some actions add pagination as part of their request body.
		if r.method == http.MethodPost {
			p, ok := r.payload.(*SearchRequest) // for /searches
			if ok {
				p.Pagination = &pagination{
					Page:    r.page,
					PerPage: r.perPage,
				}
				r.payload = p
			}

		} else {
			v := url.Values{}
			v.Set("page", strconv.Itoa(r.page))
			v.Add("per_page", strconv.Itoa(r.perPage))
			r.path += "?" + v.Encode()
		}
	}

}
