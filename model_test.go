package clarifai

import (
	"net/http"
	"testing"
	"time"
)

func TestSession_Predict(t *testing.T) {

	mux.HandleFunc("/"+apiVersion+"/models/"+PublicModelGeneral+"/outputs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")

		PrintMock(t, w, "resp/ok_predict_1img.json")
	})

	sess.tokenExpiration = time.Now().Second() + 3600 // imitate existence of non-expired token

	r := InitInputs()
	_ = r.AddImageInput(NewImageFromURL("https://samples.clarifai.com/metro-north.jpg"), "")
	resp, err := sess.Predict(r).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Outputs: []*Output{
			{
				ID: "cf0e878cd2304d888caa2bcb69a77f56",
				Status: &ServiceStatus{
					Code:        10000,
					Description: "Ok",
				},
				CreatedAt: "2016-11-29T03:15:05Z",
				Model: &Model{
					Name:      "general-v1.3",
					ID:        "aaa03c23b3724a16a56b629203edc62c",
					CreatedAt: "2016-03-09T17:11:39Z",
					AppID:     "",
					OutputInfo: &OutputInfo{
						Message: "Show output_info with: GET /models/{model_id}/output_info",
						Type:    "concept",
					},
					ModelVersion: &ModelVersion{
						ID:        "aa9ca48295b37401f8af92ad1af0d91d",
						CreatedAt: "2016-07-13T01:19:12Z",
						Status: &ServiceStatus{
							Code:        21100,
							Description: "Model trained successfully",
						},
					},
				},
				Input: &Input{
					Data: &Image{
						Properties: &ImageProperties{
							URL: "https://samples.clarifai.com/metro-north.jpg",
						},
					},
					ID: "cf0e878cd2304d888caa2bcb69a77f56",
				},
				Data: &OutputData{
					Concepts: &OutputConcepts{
						{
							ID:    "ai_HLmqFqBf",
							Name:  "train",
							Value: 0.9989112,
						},
						{
							ID:    "ai_fvlBqXZR",
							Name:  "railway",
							Value: 0.9975532,
						},
					},
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}
