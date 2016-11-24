package clarifai

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewPredictService(t *testing.T) {
	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	actual := reflect.TypeOf(svc.Session).String()
	expected := "*clarifai.Session"

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}

	actual = svc.ModelID
	expected = PublicModelGeneral

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}

	actual3 := len(svc.GetInputObject().Inputs)
	expected3 := 0

	if actual3 != expected3 {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestPredictService_GetPredictions_Fail_NoInputs(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)
	res, err := svc.GetPredictions()

	expected := &PredictResponse{}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Actual: %v, expected: %v", res, expected)
	}

	if err != ErrNoInputs {
		t.Error("Should return " + ErrNoInputs.Error())
	}
}

func TestPredictService_AddInput_Limit(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	var err error

	for i := 0; i < 128; i++ {
		err = svc.AddInput(&Input{})
		if err != nil {
			t.Fatal("Should return no errors")
		}
	}

	err = svc.AddInput(&Input{})
	if err != ErrInputLimitReached {
		t.Error("Should return " + ErrInputLimitReached.Error())
	}
}

func TestPredictService_AddInput_Success(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	err := svc.AddInput(&Input{})
	if err != nil {
		t.Errorf("Should have no errors, but got %+v", err)
	}

	actual := svc.GetInputObject().GetInputsQty()
	expected := 1

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestNewPredictService_SetModel(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	svc.SetModel(PublicModelFood)

	actual := svc.ModelID
	expected := PublicModelFood

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestNewPredictService_GetModel(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	svc.SetModel(PublicModelFood)

	actual := svc.GetModel()
	expected := PublicModelFood

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestPredictService_GetPredictions(t *testing.T) {

	mux.HandleFunc("/"+apiVersion+"/models/"+PublicModelGeneral+"/outputs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")

		PrintMock(t, w, "resp/ok_predict_1img.json")
	})

	sess.TokenExpiration = time.Now().Second() + 3600 // imitate existence of non-expired token

	svc := NewPredictService(sess)
	_ = svc.AddInput(ImageInputFromURL("https://samples.clarifai.com/metro-north.jpg", nil))

	resp, err := svc.GetPredictions()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &PredictResponse{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Outputs: []*Output{
			{
				ID: "c63d80ec8e5c42259e776e776f6ccd09",
				Status: &ServiceStatus{
					Code:        10000,
					Description: "Ok",
				},
				CreatedAt: "2016-11-24T11:41:44Z",
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
					ID: "c63d80ec8e5c42259e776e776f6ccd09",
					Data: &InputData{
						Image: &ImageData{
							URL: "https://samples.clarifai.com/metro-north.jpg",
						},
					},
				},
				Data: &OutputData{
					Concepts: &OutputConcepts{
						{
							ID:    "ai_l8TKp2h5",
							Name:  "people",
							AppID: "",
							Value: 0.99921584,
						},
						{
							ID:    "ai_VPmHr5bm",
							Name:  "adult",
							AppID: "",
							Value: 0.9947057,
						},
					},
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}
