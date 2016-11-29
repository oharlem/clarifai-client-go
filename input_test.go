package clarifai

import (
	"net/http"
	"testing"
	"time"
)

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

func TestImageInputFromURL(t *testing.T) {
	url := "https://samples.clarifai.com/metro-north.jpg"

	i := NewImageFromURL(url)

	if i.Properties.URL != url {
		t.Errorf("Actual: %v, expected: %v", i.Properties.URL, url)
	}
}

func TestInputService_ListAllInputs(t *testing.T) {

	mux.HandleFunc("/"+apiVersion+"/inputs", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)

		w.Header().Set("Content-Type", "application/json")

		PrintMock(t, w, "resp/ok_inputs.json")
	})

	sess.TokenExpiration = time.Now().Second() + 3600 // imitate existence of non-expired token

	r := NewRequest(sess)
	resp, err := sess.ListAllInputs(r)
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
