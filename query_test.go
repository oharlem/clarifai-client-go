package clarifai

import (
	"testing"
)

func TestQueryOutput_AddConcept(t *testing.T) {

	q := &QueryOutput{}
	q.AddConcept("foo", true)

	if len(q.Data.Concepts) == 0 {
		t.Fatal("Concept not added!")
	}

	actual, ok := q.Data.Concepts[0]["name"]
	if !ok {
		t.Fatal("Invalid concept name!")
	}
	expected := "foo"

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}

	actual2, ok := q.Data.Concepts[0]["value"]
	if !ok {
		t.Fatal("Invalid concept value!")
	}
	expected2 := true

	if actual2 != expected2 {
		t.Fatalf("Actual: %v, expected: %v", actual2, expected2)
	}
}
