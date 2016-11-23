package clarifai

import (
	"reflect"
	"testing"
)

const (
	mockClientID     = "foo"
	mockClientSecret = "bar"
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

	actual := getAPIHost("foo")
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
