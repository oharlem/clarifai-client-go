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

	if sess.clientID != mockClientID {
		t.Errorf("Actual: %v, expected: %v", sess.clientID, mockClientID)
	}

	if sess.clientSecret != mockClientSecret {
		t.Errorf("Actual: %v, expected: %v", sess.clientSecret, mockClientSecret)
	}

	if sess.accessToken != "" {
		t.Errorf("Actual: %v, expected: %v", sess.accessToken, "to be empty")
	}
}

func TestBuildURI(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)

	actual := sess.buildURI("foo")
	expected := apiHost + "/" + apiVersion + "/" + "foo"

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestAuthResponseValidation_Success(t *testing.T) {

	resp := &AuthResponse{
		AccessToken: "foo",
	}

	err := authResponseValidation(resp)
	if err != nil {
		t.Errorf("Should return no error, but got %v", err)
	}
}

func TestAuthResponseValidation_Fail_No_Token(t *testing.T) {

	actual := authResponseValidation(&AuthResponse{})
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

	startingTime, expirationTime := sess.setTokenExpiration(resp.ExpiresIn)

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
	sess.setTokenExpiration(resp.ExpiresIn)

	actual := sess.isTokenExpired()
	expected := false

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestAuthResponseValidation_IsTokenExpired_True(t *testing.T) {

	sess := NewSession(mockClientID, mockClientSecret)
	sess.tokenExpiration = 0

	actual := sess.isTokenExpired()
	expected := true

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestSession_Connect(t *testing.T) {

	mockRoute(t, "token", "resp/ok_auth.json")

	err := sess.Connect()
	if err != nil {
		t.Fatalf("Should have no errors, but got %v", err)
	}

	actual := sess.accessToken
	expected := "bCGdwie3gIJoRISG5Ejz2Je57inNTj"

	if expected != actual {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}

func TestApiKeyAuth(t *testing.T) {
	mux.HandleFunc("/"+apiVersion+"/key-test", func(w http.ResponseWriter, r *http.Request) {
		actual := r.Header.Get("Authorization")
		expected := "Key test_api_key"
		if expected != actual {
			t.Errorf("Actual: %v, expected: %v", actual, expected)
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok":true}`))
	})

	app := NewApp("test_api_key")
	app.host = ts.URL
	app.HTTPCall("GET", "key-test", nil)
}
