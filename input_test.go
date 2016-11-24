package clarifai

import (
	"net/http"
	"testing"
	"time"
)

func TestImageInputFromPath(t *testing.T) {
	path := "mocks/test_image.jpg"

	i, err := ImageInputFromPath(path)

	if err != nil {
		t.Fatalf("Should have no errors, but got %+v", err)
	}

	if i.Data.Image.Base64 != TestImageBase64 {
		t.Errorf("Actual: %v, expected: %v", i.Data.Image.Base64, TestImageBase64)
	}
}

func TestValidateLocalFile_Fail(t *testing.T) {
	path := "mocks/unsupported_mime_type_gif.gif"
	expected := ErrUnsupportedMimeType

	_, err := ImageInputFromPath(path)
	if err != expected {
		t.Errorf("Actual: %v, expected: %v", err, expected)
	}
}

func TestImageInputFromURL(t *testing.T) {
	url := "https://samples.clarifai.com/metro-north.jpg"

	i := ImageInputFromURL(url, nil)

	if i.Data.Image.URL != url {
		t.Errorf("Actual: %v, expected: %v", i.Data.Image.URL, url)
	}
}

func TestInputService_ListAllInputs(t *testing.T) {

	mux.HandleFunc("/"+apiVersion+"/inputs", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)

		w.Header().Set("Content-Type", "application/json")

		PrintMock(t, w, "resp/ok_inputs.json")
	})

	sess.TokenExpiration = time.Now().Second() + 3600 // imitate existence of non-expired token

	svc := NewInputService(sess)
	resp, err := svc.ListAllInputs()

	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &InputResponse{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Inputs: []*Input{
			{
				ID:        "ce8524a1191d4b47816d07a0f4d06b36",
				CreatedAt: "2016-11-21T06:10:09Z",
				Data: &InputData{
					Image: &ImageData{
						URL: "https://samples.clarifai.com/metro-north.jpg",
					},
				},
				Status: &ServiceStatus{
					Code:        30000,
					Description: "Download complete",
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}
