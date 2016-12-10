package clarifai

import (
	"reflect"
	"testing"
)

func TestImageInputFromURL(t *testing.T) {
	url := "https://samples.clarifai.com/metro-north.jpg"

	i := NewImageFromURL(url)

	if i.Properties.URL != url {
		t.Errorf("Actual: %v, expected: %v", i.Properties.URL, url)
	}
}

func TestImageInputFromPath(t *testing.T) {
	path := "mocks/test_image.jpg"

	i, err := NewImageFromFile(path)

	if err != nil {
		t.Fatalf("Should have no errors, but got %+v", err)
	}

	if i.Properties.Base64 != TestImageBase64 {
		t.Errorf("Actual: %v, expected: %v", i.Properties.Base64, TestImageBase64)
	}
}

func TestValidateLocalFile_Fail(t *testing.T) {
	path := "mocks/unsupported_mime_type_gif.gif"
	expected := ErrUnsupportedMimeType

	_, err := NewImageFromFile(path)
	if err != expected {
		t.Errorf("Actual: %v, expected: %v", err, expected)
	}
}

func TestImage_AddCrop(t *testing.T) {
	url := "https://samples.clarifai.com/metro-north.jpg"

	i := NewImageFromURL(url)
	i.AddCrop(0.2, 0.4, 0.3, 0.6)

	expected := []float32{0.2, 0.4, 0.3, 0.6}
	actual := i.Properties.Crop

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestImage_AddConcept(t *testing.T) {

	i := NewImageFromURL("https://samples.clarifai.com/metro-north.jpg")
	i.AddConcept("foo", true)

	expected := map[string]interface{}{
		"id":    "foo",
		"value": true,
	}
	actual := i.Concepts[0]

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestImage_AddConcepts(t *testing.T) {

	i := NewImageFromURL("https://samples.clarifai.com/metro-north.jpg")
	i.AddConcepts([]string{"foo", "bar"})

	if len(i.Concepts) != 2 {
		t.Errorf("Actual: %v, expected: %v", len(i.Concepts), 2)
	}

	expected1 := map[string]interface{}{
		"id":    "foo",
		"value": true,
	}
	if !reflect.DeepEqual(i.Concepts[0], expected1) {
		t.Errorf("Actual: %v, expected: %v", i.Concepts[0], expected1)
	}

	expected2 := map[string]interface{}{
		"id":    "bar",
		"value": true,
	}
	if !reflect.DeepEqual(i.Concepts[1], expected2) {
		t.Errorf("Actual: %v, expected: %v", i.Concepts[1], expected2)
	}
}

func TestImage_AddMetadata(t *testing.T) {

	i := NewImageFromURL("https://samples.clarifai.com/metro-north.jpg")
	i.AddMetadata(map[string]string{
		"event_type": "vacation",
	})

	expected := map[string]string{
		"event_type": "vacation",
	}
	actual := i.Metadata

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestImage_AllowDuplicates(t *testing.T) {

	i := NewImageFromURL("https://samples.clarifai.com/metro-north.jpg")
	i.AllowDuplicates()

	if i.Properties.AllowDuplicateURL != true {
		t.Errorf("Actual: %v, expected: %v", i.Properties.AllowDuplicateURL, true)
	}
}
