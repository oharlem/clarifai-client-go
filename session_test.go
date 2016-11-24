package clarifai

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewSession(t *testing.T) {
	sess := NewSession(mockClientID, mockClientSecret)

	actual := reflect.TypeOf(sess).String()
	expected := "*clarifai.Session"

	if actual != expected {
		t.Fatalf("Actual: %v, expected: %v", actual, expected)
	}

	if sess.ClientID != mockClientID {
		t.Errorf("Actual: %v, expected: %v", sess.ClientID, mockClientID)
	}

	if sess.ClientSecret != mockClientSecret {
		t.Errorf("Actual: %v, expected: %v", sess.ClientSecret, mockClientSecret)
	}

	if sess.AccessToken != "" {
		t.Errorf("Actual: %v, expected: %v", sess.AccessToken, "to be empty")
	}
}

func TestGetAPIHost(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)

	actual := sess.GetAPIHost("foo")
	expected := apiHost + "/" + apiVersion + "/" + "foo"

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestSetAccessToken(t *testing.T) {
	sess := NewSession(mockClientID, mockClientSecret)

	sess.setAccessToken("foo")

	actual := sess.AccessToken
	expected := "foo"

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestGetAccessToken(t *testing.T) {
	sess := NewSession(mockClientID, mockClientSecret)

	sess.AccessToken = "foo"

	actual := sess.GetAccessToken()
	expected := "foo"

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestAuthResponseValidation_Success(t *testing.T) {

	resp := &AuthResponse{
		AccessToken: "foo",
	}

	err := AuthResponseValidation(resp)
	if err != nil {
		t.Errorf("Should return no error, but got %v", err)
	}
}

func TestAuthResponseValidation_Fail_No_Token(t *testing.T) {

	actual := AuthResponseValidation(&AuthResponse{})
	expected := ErrNoAuthenticationToken

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestAuthResponseValidation_SetTokenExpiration(t *testing.T) {

	expiresIn := 100

	sess := NewSession(mockClientID, mockClientSecret)

	resp := &AuthResponse{
		ExpiresIn: expiresIn,
	}

	startingTime, expirationTime := sess.SetTokenExpiration(resp.ExpiresIn)

	if expirationTime != startingTime+expiresIn {
		t.Errorf("Actual: %v, expected: %v", expirationTime, startingTime+expiresIn)
	}
}

func TestAuthResponseValidation_IsTokenExpired_False(t *testing.T) {

	expiresIn := 100

	sess := NewSession(mockClientID, mockClientSecret)

	resp := &AuthResponse{
		ExpiresIn: expiresIn,
	}
	sess.SetTokenExpiration(resp.ExpiresIn)

	actual := sess.IsTokenExpired()
	expected := false

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestAuthResponseValidation_IsTokenExpired_True(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	sess.TokenExpiration = 0

	actual := sess.IsTokenExpired()
	expected := true

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestSession_Connect(t *testing.T) {

	mux.HandleFunc("/"+apiVersion+"/token", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)

		w.Header().Set("Content-Type", "application/json")

		PrintMock(t, w, "resp/ok_auth.json")
	})

	err := sess.Connect()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	actual := sess.GetAccessToken()
	expected := "bCGdwie3gIJoRISG5Ejz2Je57inNTj"

	if expected != actual {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}
