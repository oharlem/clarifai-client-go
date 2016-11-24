package clarifai

import "errors"

var (
	ErrNoAuthenticationToken = errors.New("No authentication token returned!")
	ErrNoInputs              = errors.New("No inputs specified!")
	ErrInputLimitReached     = errors.New("Reached maximum number of allowed inputs!")
	ErrUnsupportedMimeType   = errors.New("Input with an unsupported mime type provided!")
)
