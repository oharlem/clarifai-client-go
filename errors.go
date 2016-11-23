package clarifai

import "errors"

var (
	ErrNoInputs          = errors.New("No inputs specified!")
	ErrInputLimitReached = errors.New("Reached maximum number of allowed inputs!")
)
