package clarifai

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {

	s := NewSession(mockClientID, mockClientSecret)

	r := NewRequest(s, http.MethodGet, "foo")

	if r.method != http.MethodGet {
		t.Errorf("Invalid method | Actual: %v, expected: %v", r.method, http.MethodGet)
	}

	if r.page != defaultPage {
		t.Errorf("Invalid page | Actual: %v, expected: %v", r.page, defaultPage)
	}

	if r.perPage != defaultItemsPerPageQty {
		t.Errorf("Invalid perPage | Actual: %v, expected: %v", r.perPage, defaultItemsPerPageQty)
	}

	if r.path != "foo" {
		t.Errorf("Invalid path | Actual: %v, expected: %v", r.path, "foo")
	}

	if r.payload != nil {
		t.Errorf("Invalid payload | Actual: %v, expected: %v", r.payload, nil)
	}

	if !reflect.DeepEqual(r.session, s) {
		t.Errorf("Invalid session | Actual: %v, expected: %v", r.session, s)
	}
}

func TestRequest_WithPagination(t *testing.T) {

	r := NewRequest(sess, http.MethodPost, "foo")
	r.WithPagination(1, 5)

	if r.page != 1 {
		t.Errorf("Actual: %v, expected: %v", r.page, 1)
	}

	if r.perPage != 5 {
		t.Errorf("Actual: %v, expected: %v", r.perPage, 5)
	}
}

func TestRequest_addPagination_Get(t *testing.T) {

	r := NewRequest(sess, http.MethodGet, "foo")
	r.WithPagination(1, 5)
	r.addPagination()

	expected := "foo?page=1&per_page=5"

	if r.path != expected {
		t.Errorf("Actual: %v, expected: %v", r.path, expected)
	}
}

func TestRequest_addPagination_Post(t *testing.T) {

	// addPagination adds pagination property to search requests only: "searches", "concepts/searches"
	r := NewRequest(sess, http.MethodPost, "searches")
	r.WithPagination(1, 5)

	r.payload = NewAndSearchQuery()

	r.addPagination()

	p, ok := r.payload.(*SearchRequest)
	if !ok {
		t.Fatalf("Actual: %+v, expected: %+v", reflect.TypeOf(p).String(), "*clarifai.SearchRequest")
	}

	if p.Pagination == nil {
		t.Fatal("Pagination property should not be nil")
	}

	if p.Pagination.Page != 1 {
		t.Errorf("Actual: %v, expected: %v", p.Pagination.Page, 1)
	}

	if p.Pagination.PerPage != 5 {
		t.Errorf("Actual: %v, expected: %v", p.Pagination.PerPage, 5)
	}
}
