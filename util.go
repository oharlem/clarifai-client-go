package clarifai

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// PP is a pretty-print helper for debugging purposes.
func PP(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

// EncBytesToBase64 Encodes []byte input into a base64 string.
func EncBytesToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
