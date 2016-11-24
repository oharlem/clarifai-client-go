package clarifai

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// PE returns prettified object info.
func PE(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("error %v\n", err)
	}
	return fmt.Sprintf("%v\n", string(b))
}

// PP prints out prettified object info.
func PP(v interface{}) {
	fmt.Print(PE(v))
}

// EncBytesToBase64 Encodes []byte input into a base64 string.
func EncBytesToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
