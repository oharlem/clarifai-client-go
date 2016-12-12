package clarifai

import (
	"reflect"
	"testing"
)

func TestNewSearchQuery(t *testing.T) {
	q := NewSearchQuery(SearchQueryTypeAnd)

	if q.Type != SearchQueryTypeAnd {
		t.Errorf("Type | Actual: %v, expected: %v", q.Type, SearchQueryTypeAnd)
	}

	if q.QueryObject == nil {
		t.Error("QueryObject | Should not be empty")
	}
}

func TestNewAndSearchQuery(t *testing.T) {
	q := NewAndSearchQuery()

	if q.Type != SearchQueryTypeAnd {
		t.Errorf("Type | Actual: %v, expected: %v", q.Type, SearchQueryTypeAnd)
	}

	if q.QueryObject == nil {
		t.Error("QueryObject | Should not be empty")
	}
}

func TestSearchRequest_WithUserConcept(t *testing.T) {

	q := NewAndSearchQuery()
	q.WithUserConcept("foo")

	expected := map[string]interface{}{
		"name":  "foo",
		"value": true,
	}

	if !reflect.DeepEqual(q.QueryObject.Ands[0].Input.Data.Concepts[0], expected) {
		t.Errorf("Input concept | Actual: %v, expected: %v", q.QueryObject.Ands[0].Input.Data.Concepts[0], expected)
	}
}

func TestSearchRequest_WithoutUserConcept(t *testing.T) {

	q := NewAndSearchQuery()
	q.WithoutUserConcept("foo")

	expected := map[string]interface{}{
		"name":  "foo",
		"value": false,
	}

	if !reflect.DeepEqual(q.QueryObject.Ands[0].Input.Data.Concepts[0], expected) {
		t.Errorf("Input concept | Actual: %v, expected: %v", q.QueryObject.Ands[0].Input.Data.Concepts[0], expected)
	}
}

func TestSearchRequest_WithAPIConcept(t *testing.T) {

	q := NewAndSearchQuery()
	q.WithAPIConcept("foo")

	expected := map[string]interface{}{
		"name":  "foo",
		"value": true,
	}

	if !reflect.DeepEqual(q.QueryObject.Ands[0].Output.Data.Concepts[0], expected) {
		t.Errorf("Output concept | Actual: %v, expected: %v", q.QueryObject.Ands[0].Output.Data.Concepts[0], expected)
	}
}

func TestSearchRequest_WithoutAPIConcept(t *testing.T) {

	q := NewAndSearchQuery()
	q.WithoutAPIConcept("foo")

	expected := map[string]interface{}{
		"name":  "foo",
		"value": false,
	}

	if !reflect.DeepEqual(q.QueryObject.Ands[0].Output.Data.Concepts[0], expected) {
		t.Errorf("Output concept | Actual: %v, expected: %v", q.QueryObject.Ands[0].Output.Data.Concepts[0], expected)
	}

}

func TestSearchRequest_WithImage(t *testing.T) {

	q := NewAndSearchQuery()
	i := NewImageFromURL("https://samples.clarifai.com/metro-north.jpg")
	q.WithImage(i)

	if q.QueryObject.Ands[0].Output.Input.Data.Properties.URL != "https://samples.clarifai.com/metro-north.jpg" {
		t.Errorf("WithImage | Actual: %v, expected: %v", q.QueryObject.Ands[0].Output.Data.Concepts[0], "https://samples.clarifai.com/metro-north.jpg")
	}
}

func TestSearchRequest_WithMetadata(t *testing.T) {

	q := NewAndSearchQuery()
	expected := map[string]interface{}{
		"event_type": "vacation",
	}
	q.WithMetadata(expected)

	if !reflect.DeepEqual(q.QueryObject.Ands[0].Input.Data.Metadata, expected) {
		t.Errorf("WithMetadata | Actual: %v, expected: %v", q.QueryObject.Ands[0].Input.Data.Metadata, expected)
	}
}
