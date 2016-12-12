package clarifai

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
)

type Image struct {
	Concepts   []map[string]interface{} `json:"concepts,omitempty"`
	Metadata   interface{}              `json:"metadata,omitempty"`
	Properties *ImageProperties         `json:"image, omitempty"`
}

type ImageData struct {
	Concepts   []map[string]interface{} `json:"concepts,omitempty"`
	Metadata   interface{}              `json:"metadata,omitempty"`
	Properties *ImageProperties         `json:"image, omitempty"`
}

type ImageProperties struct {
	AllowDuplicateURL bool      `json:"allow_duplicate_url,omitempty"`
	Base64            string    `json:"base64,omitempty"`
	URL               string    `json:"url,omitempty"`
	Crop              []float32 `json:"crop,omitempty"`
}

// SupportedMimeTypes is a map of supported image types
// as per https://developer-preview.clarifai.com/guide/#supported-types
var SupportedMimeTypes map[string]struct{}

func init() {
	SupportedMimeTypes = map[string]struct{}{
		"image/bmp":  struct{}{},
		"image/jpeg": struct{}{},
		"image/png":  struct{}{},
		"image/tiff": struct{}{},
	}
}

// NewImageFromURL instantiates a new image based on URL.
func NewImageFromURL(url string) *Image {

	return &Image{
		Properties: &ImageProperties{
			URL: url,
		},
	}
}

// NewImageFromURL instantiates a new image from a local file.
func NewImageFromFile(path string) (*Image, error) {

	base64Str, err := addFromBase64(path)
	if err != nil {
		return &Image{}, err
	}

	return &Image{
		Properties: &ImageProperties{
			Base64: base64Str,
		},
	}, nil
}

// AllowDuplicates enables image duplicates.
func (i *Image) AllowDuplicates() {
	if i.Properties == nil {
		i.Properties = &ImageProperties{}
	}
	i.Properties.AllowDuplicateURL = true
}

// AddMetadata adds an image metadata.
func (i *Image) AddMetadata(m interface{}) {
	i.Metadata = m
}

// AddCrop adds an image metadata.
func (i *Image) AddCrop(args ...float32) {

	if i.Properties == nil {
		i.Properties = &ImageProperties{}
	}

	for _, v := range args {
		i.Properties.Crop = append(i.Properties.Crop, v)
	}
}

// AddConcept adds an image concept.
func (i *Image) AddConcept(id string, value interface{}) {

	i.Concepts = append(i.Concepts, map[string]interface{}{
		"id":    id,
		"value": value,
	})
}

// AddConcepts adds a list of concepts to an image.
func (i *Image) AddConcepts(c []string) {
	for _, v := range c {
		i.AddConcept(v, true)
	}
}

// addFromBase64 reads a local image, validates it and returns it as a base64 string.
func addFromBase64(filename string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	valErr := validateLocalFile(data)
	if valErr != nil {
		return "", valErr
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// validateLocalFile validates contents of the locally provided image file.
func validateLocalFile(data []byte) error {

	mimeType := http.DetectContentType(data)
	_, ok := SupportedMimeTypes[mimeType]
	if !ok {
		return ErrUnsupportedMimeType
	}

	return nil
}
