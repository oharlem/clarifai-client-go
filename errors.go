package clarifai

import "errors"

var (
	ErrNoAuthenticationToken = errors.New("No authentication token returned!")
	ErrInputLimitReached     = errors.New("Reached maximum number of allowed inputs!")
	ErrUnsupportedMimeType   = errors.New("Image input with an unsupported mime type provided!")
)
