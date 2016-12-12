package clarifai

import (
	"net/http"
	"reflect"
	"testing"
)

func TestInitInputs(t *testing.T) {
	i := InitInputs()

	if reflect.TypeOf(i).String() != "*clarifai.Inputs" {
		t.Fatalf("Actual: %v, expected: *clarifai.Inputs", reflect.TypeOf(i))
	}
	actual := i.modelID
	expected := PublicModelGeneral

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestInputs_AddInput(t *testing.T) {

	data := InitInputs()
	i := NewImageFromURL("https://samples.clarifai.com/travel.jpg")

	_ = data.AddInput(i, "travel-1")

	expected := "https://samples.clarifai.com/travel.jpg"
	actual := data.Inputs[0].Data.Properties.URL

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestInputs_AddInput_Limit(t *testing.T) {

	data := InitInputs()
	i := NewImageFromURL("https://samples.clarifai.com/travel.jpg")

	for j := 0; j < InputLimit; j++ {
		_ = data.AddInput(i, "travel-1")
	}

	// This element should be out of limit.
	err := data.AddInput(i, "travel-1")
	if err != ErrInputLimitReached {
		t.Fatalf("Expected limit error \"%v\", but got: %v", ErrInputLimitReached, err)
	}
}

func TestInputs_SetModel(t *testing.T) {

	i := InitInputs()
	i.SetModel(PublicModelFood)

	actual := i.modelID
	expected := PublicModelFood

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestInput_AddConcept(t *testing.T) {

	i := &Input{}
	i.AddConcept("foo", true)

	expected := map[string]interface{}{
		"name":  "foo",
		"value": true,
	}
	actual := i.Data.Concepts[0]

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestInput_SetMetadata(t *testing.T) {
	i := &Input{}

	i.SetMetadata(map[string]interface{}{
		"event_type": "vacation",
	})

	actual := i.Data.Metadata
	expected := map[string]interface{}{
		"event_type": "vacation",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestSession_GetAllInputs(t *testing.T) {

	mockRoute(t, "inputs", "resp/ok_inputs.json")

	resp, err := sess.GetAllInputs().Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Inputs: []*Input{
			{
				Data: &Image{
					Properties: &ImageProperties{
						URL: "https://samples.clarifai.com/metro-north.jpg",
					},
				},
				ID:        "ce8524a1191d4b47816d07a0f4d06b36",
				CreatedAt: "2016-11-21T06:10:09Z",
				Status: &ServiceStatus{
					Code:        30000,
					Description: "Download complete",
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_AddInputs(t *testing.T) {

	data := InitInputs()
	i := NewImageFromURL("https://samples.clarifai.com/travel.jpg")
	_ = data.AddInput(i, "travel-1")

	resp := sess.AddInputs(data)
	if reflect.TypeOf(resp).String() != "*clarifai.Request" {
		t.Fatalf("Actual: %v, expected: %v", reflect.TypeOf(resp).String(), "*clarifai.Request")
	}

	if reflect.TypeOf(resp.session).String() != "*clarifai.Session" {
		t.Fatalf("Actual: %v, expected: %v", reflect.TypeOf(resp.session).String(), "*clarifai.Session")
	}

	if resp.method != http.MethodPost {
		t.Fatal("Method should be POST.")
	}

	if resp.path != "inputs" {
		t.Fatalf("Actual: %v, expected: %v", resp.path, "inputs")
	}

	if resp.payload == nil {
		t.Fatal("Payload should not be nil.")
	}
}
