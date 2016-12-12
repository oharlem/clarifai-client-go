package clarifai

import (
	"testing"
)

func TestSession_Predict(t *testing.T) {

	mockRoute(t, "models/"+PublicModelGeneral+"/outputs", "resp/ok_predict_1img.json")

	r := InitInputs()
	_ = r.AddInput(NewImageFromURL("https://samples.clarifai.com/metro-north.jpg"), "")
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
					Name:      String("general-v1.3"),
					ID:        String("aaa03c23b3724a16a56b629203edc62c"),
					CreatedAt: String("2016-03-09T17:11:39Z"),
					AppID:     String(""),
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
					Concepts: []*OutputConcept{
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

func TestSession_CreateModel(t *testing.T) {

	mockRoute(t, "models", "resp/ok_10000_create_model.json")

	opt := NewModelOptions()
	opt.ID = "test-id-1"

	resp, err := sess.CreateModel("test-model", opt).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Model: &Model{
			Name:      String("test-model"),
			ID:        String("test-id-1"),
			CreatedAt: String("2016-12-12T03:21:51Z"),
			AppID:     String("c3915e768bf44e1eb469483642a664ef"),
			OutputInfo: &OutputInfo{
				Message: "Show output_info with: GET /models/{model_id}/output_info",
				Type:    "concept",
				OutputConfig: &OutputConfig{
					ConceptsMutuallyExclusive: false,
					ClosedEnvironment:         false,
				},
			},
			ModelVersion: &ModelVersion{
				ID:        "f9320f4aa9be46bbaa1384fc12f61c1d",
				CreatedAt: "2016-12-12T03:21:51Z",
				Status: &ServiceStatus{
					Code:        21102,
					Description: "Model not yet trained",
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_GetModels(t *testing.T) {

	serverReset()
	mockRoute(t, "models", "resp/ok_10000_get_models.json")

	resp, err := sess.GetModels().Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Models: []*Model{&Model{
			Name:      String("general-v1.3"),
			ID:        String("eab1fd01a5544225b32d5d2937e05041"),
			CreatedAt: String("2016-11-24T18:15:43Z"),
			AppID:     String("c3915e768bf44e1eb469483642a664ef"),
			OutputInfo: &OutputInfo{
				Message: "Show output_info with: GET /models/{model_id}/output_info",
				Type:    "concept",
			},
			ModelVersion: &ModelVersion{
				ID:        "d88847bb75514fceaf74bf36606d1343",
				CreatedAt: "2016-12-11T00:54:39Z",
				Status: &ServiceStatus{
					Code:        21111,
					Description: "Model training had no positive examples.",
				},
			},
		}},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_GetModel(t *testing.T) {

	serverReset()

	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3

	mockRoute(t, "models/"+modelID, "resp/ok_10000_get_model.json")

	resp, err := sess.GetModel(modelID).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Model: &Model{
			Name:      String("general-v1.3"),
			ID:        String("eab1fd01a5544225b32d5d2937e05041"),
			CreatedAt: String("2016-11-24T18:15:43Z"),
			AppID:     String("c3915e768bf44e1eb469483642a664ef"),
			OutputInfo: &OutputInfo{
				Message: "Show output_info with: GET /models/{model_id}/output_info",
				Type:    "concept",
			},
			ModelVersion: &ModelVersion{
				ID:        "d88847bb75514fceaf74bf36606d1343",
				CreatedAt: "2016-12-11T00:54:39Z",
				Status: &ServiceStatus{
					Code:        21111,
					Description: "Model training had no positive examples.",
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_GetModelOutput(t *testing.T) {

	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3

	mockRoute(t, "models/"+modelID+"/output_info", "resp/ok_10000_get_model_output_info.json")

	resp, err := sess.GetModelOutput(modelID).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Model: &Model{
			Name:      String("general-v1.3"),
			ID:        String("eab1fd01a5544225b32d5d2937e05041"),
			CreatedAt: String("2016-11-24T18:15:43Z"),
			AppID:     String("c3915e768bf44e1eb469483642a664ef"),
			OutputInfo: &OutputInfo{
				Message: "",
				Type:    "concept",
				OutputConfig: &OutputConfig{
					ConceptsMutuallyExclusive: false,
					ClosedEnvironment:         false,
				},
				OutputData: &OutputData{
					Concepts: []*OutputConcept{
						{
							AppID:     String("c3915e768bf44e1eb469483642a664ef"),
							CreatedAt: "2016-11-27T23:10:05Z",
							ID:        "album",
							Name:      "album",
							UpdatedAt: "2016-11-27T23:10:05Z",
						},
						{
							AppID:     String("c3915e768bf44e1eb469483642a664ef"),
							CreatedAt: "2016-11-28T02:32:36Z",
							ID:        "test",
							Name:      "test",
							UpdatedAt: "2016-11-28T02:32:36Z",
						},
					},
				},
			},
			ModelVersion: &ModelVersion{
				ID:        "d88847bb75514fceaf74bf36606d1343",
				CreatedAt: "2016-12-11T00:54:39Z",
				Status: &ServiceStatus{
					Code:        21111,
					Description: "Model training had no positive examples.",
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_GetModelVersion(t *testing.T) {
	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3
	versionID := "d88847bb75514fceaf74bf36606d1343"

	mockRoute(t, "models/"+modelID+"/versions/"+versionID, "resp/ok_10000_get_model_version.json")

	resp, err := sess.GetModelVersion(modelID, versionID).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		ModelVersion: &ModelVersion{
			ID:        "d88847bb75514fceaf74bf36606d1343",
			CreatedAt: "2016-12-11T00:54:39Z",
			Status: &ServiceStatus{
				Code:        21111,
				Description: "Model training had no positive examples.",
			},
		},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_GetModelVersions(t *testing.T) {
	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3

	mockRoute(t, "models/"+modelID+"/versions", "resp/ok_10000_get_model_versions.json")

	resp, err := sess.GetModelVersions(modelID).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		ModelVersions: []*ModelVersion{
			{
				ID:        "d88847bb75514fceaf74bf36606d1343",
				CreatedAt: "2016-12-11T00:54:39Z",
				Status: &ServiceStatus{
					Code:        21111,
					Description: "Model training had no positive examples.",
				},
			},
			{
				ID:        "c85c99ba18aa4b0da0a597829b13ca37",
				CreatedAt: "2016-12-10T04:00:15Z",
				Status: &ServiceStatus{
					Code:        21111,
					Description: "Model training had no positive examples.",
				},
			},
		},
	}

	CompareStructs(t, expected, resp)
}

func TestSession_GetModelInputs(t *testing.T) {
	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3

	mockRoute(t, "models/"+modelID+"/inputs", "resp/ok_10000_get_model_inputs.json")

	resp, err := sess.GetModelInputs(modelID).Do()
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
					Concepts: []map[string]interface{}{
						{
							"app_id": "c3915e768bf44e1eb469483642a664ef",
							"id":     "album",
							"name":   "album",
							"value":  1.0,
						},
						{
							"app_id": "c3915e768bf44e1eb469483642a664ef",
							"id":     "vacation",
							"name":   "vacation",
							"value":  1.0,
						},
					},
					Properties: &ImageProperties{
						Crop: []float32{0.2, 0.4, 0.3, 0.6},
						URL:  "https://s3.amazonaws.com/clarifai-api/img/prod/c3915e768bf44e1eb469483642a664ef/f58595769e06422287a090aeef824eb5.jpeg",
					},
				},
				ID:        "travel-1",
				CreatedAt: "2016-12-09T05:23:16Z",
			},
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_GetModelVersionInputs(t *testing.T) {
	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3
	versionID := "d88847bb75514fceaf74bf36606d1343"

	mockRoute(t, "models/"+modelID+"/versions/"+versionID+"/inputs", "resp/ok_10000_get_model_version_inputs.json")

	resp, err := sess.GetModelVersionInputs(modelID, versionID).Do()
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
					Concepts: []map[string]interface{}{
						{
							"app_id": "c3915e768bf44e1eb469483642a664ef",
							"id":     "album",
							"name":   "album",
							"value":  1.0,
						},
						{
							"app_id": "c3915e768bf44e1eb469483642a664ef",
							"id":     "vacation",
							"name":   "vacation",
							"value":  1.0,
						},
					},
					Properties: &ImageProperties{
						Crop: []float32{0.2, 0.4, 0.3, 0.6},
						URL:  "https://foo.com/bar.jpeg",
					},
				},
				ID:        "travel-1",
				CreatedAt: "2016-12-09T05:23:16Z",
			},
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_DeleteModelVersion(t *testing.T) {
	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3
	versionID := "d88847bb75514fceaf74bf36606d1343"

	serverReset()

	mockRoute(t, "models/"+modelID+"/versions/"+versionID, "resp/ok_10000_delete_model_version.json")

	resp, err := sess.DeleteModelVersion(modelID, versionID).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_DeleteModel(t *testing.T) {
	modelID := "eab1fd01a5544225b32d5d2937e05041" // general-1.3

	serverReset()

	mockRoute(t, "models/"+modelID, "resp/ok_10000_delete_model_version.json")

	resp, err := sess.DeleteModel(modelID).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_DeleteAllModels(t *testing.T) {

	mockRoute(t, "models", "resp/ok_10000_delete_all_models.json")

	resp, err := sess.DeleteAllModels().Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_TrainModel(t *testing.T) {

	serverReset()

	modelID := "music-model-id-1"
	mockRoute(t, "models/"+modelID+"/versions", "resp/ok_10000_train_model.json")

	resp, err := sess.TrainModel("music-model-id-1").Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Model: &Model{
			Name:      String("music-model"),
			ID:        String("music-model-id-1"),
			CreatedAt: String("2016-12-12T04:17:47Z"),
			AppID:     String("c3915e768bf44e1eb469483642a664ef"),
			OutputInfo: &OutputInfo{
				Message: "Show output_info with: GET /models/{model_id}/output_info",
				Type:    "concept",
				OutputConfig: &OutputConfig{
					ConceptsMutuallyExclusive: false,
					ClosedEnvironment:         false,
				},
			},
			ModelVersion: &ModelVersion{
				ID:        "e28e2faba6da437dba046aadeb32266c",
				CreatedAt: "2016-12-12T04:19:09Z",
				Status: &ServiceStatus{
					Code:        21103,
					Description: "Custom model is currently in queue for training, waiting on inputs to process.",
				},
			},
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_DeleteModelConcepts(t *testing.T) {

	serverReset()

	modelID := "music-model-id-1"
	mockRoute(t, "models/"+modelID+"/output_info/data/concepts", "resp/ok_10000_delete_model_concepts.json")

	resp, err := sess.DeleteModelConcepts("music-model-id-1", []string{"Depeche Mode"}).Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Model: &Model{
			Name:      String("music-model"),
			ID:        String("music-model-id-1"),
			CreatedAt: String("2016-12-12T04:17:47Z"),
			AppID:     String("c3915e768bf44e1eb469483642a664ef"),
			OutputInfo: &OutputInfo{
				Message: "Show output_info with: GET /models/{model_id}/output_info",
				Type:    "concept",
				OutputConfig: &OutputConfig{
					ConceptsMutuallyExclusive: false,
					ClosedEnvironment:         false,
				},
			},
			ModelVersion: &ModelVersion{
				ID:        "d8e94afec63840978210fe2274c4b560",
				CreatedAt: "2016-12-12T05:05:26Z",
				Status: &ServiceStatus{
					Code:        21102,
					Description: "Model not yet trained",
				},
			},
		},
	}
	CompareStructs(t, expected, resp)
}

func TestSession_SearchModel(t *testing.T) {

	mockRoute(t, "models/searches", "resp/ok_10000_search_model.json")

	resp, err := sess.SearchModel("music-model", "").Do()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	expected := &Response{
		Status: &ServiceStatus{
			Code:        10000,
			Description: "Ok",
		},
		Models: []*Model{
			{
				Name:      String("music-model"),
				ID:        String("music-model-id-1"),
				CreatedAt: String("2016-12-12T04:17:47Z"),
				AppID:     String("c3915e768bf44e1eb469483642a664ef"),
				OutputInfo: &OutputInfo{
					Message: "Show output_info with: GET /models/{model_id}/output_info",
					Type:    "concept",
					OutputConfig: &OutputConfig{
						ConceptsMutuallyExclusive: false,
						ClosedEnvironment:         false,
					},
				},
				ModelVersion: &ModelVersion{
					ID:        "d8e94afec63840978210fe2274c4b560",
					CreatedAt: "2016-12-12T05:05:26Z",
					Status: &ServiceStatus{
						Code:        21102,
						Description: "Model not yet trained",
					},
				},
			},
		},
	}
	CompareStructs(t, expected, resp)
}
