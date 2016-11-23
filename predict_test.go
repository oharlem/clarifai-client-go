package clarifai

import (
	"reflect"
	"testing"
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

	actual3 := len(svc.Inputs.Inputs)
	expected3 := 0

	if actual3 != expected3 {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestPredictService_Call_Fail_NoInputs(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)
	res, err := svc.Call()

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

	actual := svc.GetInputsQty()
	expected := 1

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestPredictService_CallValidations_Success(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	_ = svc.AddInput(&Input{})

	err := svc.CallValidations()

	if err != nil {
		t.Errorf("Should have no errors, got %+v ", err)
	}
}

func TestPredictService_CallValidations_Fail(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	err := svc.CallValidations()

	if err != ErrNoInputs {
		t.Error("Should return " + ErrNoInputs.Error())
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

func TestNewPredictService_GetModelEndpoint(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	svc := NewPredictService(sess)

	actual := svc.GetModelEndpoint()
	expected := apiHost + "/" + apiVersion + "/models/" + svc.ModelID + "/outputs"

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}
}
