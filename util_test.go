package clarifai

import "testing"

func TestEncBytesToBase64(t *testing.T) {
	dataStr := "Hello, Dennis!"
	expected := "SGVsbG8sIERlbm5pcyE="

	actual := EncBytesToBase64([]byte(dataStr))

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v", actual, expected)
	}
}
