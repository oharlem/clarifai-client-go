package clarifai

import (
	"io/ioutil"
	"net/http"
	"os"
)

var AllowedMimeTypes map[string]struct{}

func init() {
	AllowedMimeTypes = map[string]struct{}{
		"image/bmp":  struct{}{},
		"image/jpeg": struct{}{},
		"image/png":  struct{}{},
		"image/tiff": struct{}{},
	}
}

type InputObject struct {
	Inputs []*Input `json:"inputs"`
}

type OutputObject struct {
	Outputs []*Output `json:"outputs"`
}

func (o *InputObject) AddInput(i *Input) {
	o.Inputs = append(o.Inputs, i)
}

func (o *InputObject) GetInputsQty() int {
	if o.Inputs == nil {
		return 0
	}
	return len(o.Inputs)
}

type Input struct {
	Data      *InputData     `json:"data,omitempty"`
	ID        string         `json:"id,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	Status    *ServiceStatus `json:"status,omitempty"`
}

type InputConcept struct {
	ID    string      `json:"id"`
	Value interface{} `json:"value,omitempty"`
}

type InputData struct {
	Concepts []*InputConcept `json:"concepts,omitempty"`
	Image    *ImageData      `json:"image,omitempty"`
	Metadata *interface{}    `json:"metadata,omitempty"`
}

func (id *InputData) AddImage(i *ImageData) {
	id.Image = i
}

func (id *InputData) AddMetadata(m *interface{}) {
	id.Metadata = m
}

type ImageData struct {
	Base64 string    `json:"base64,omitempty"`
	URL    string    `json:"url,omitempty"`
	Crop   []float32 `json:"crop,omitempty"`
}

func (i *ImageData) AddFromURL(s string) {
	i.URL = s
}

func (i *ImageData) AddFromBase64(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	validationError := ValidateLocalFile(data)
	if validationError != nil {
		return validationError
	}

	i.Base64 = EncBytesToBase64(data)

	return nil
}

// ValidateLocalFile validates contents of the locally provided image file.
func ValidateLocalFile(data []byte) error {

	mimeType := http.DetectContentType(data)
	_, ok := AllowedMimeTypes[mimeType]
	if !ok {
		return ErrUnsupportedMimeType
	}

	return nil
}

// ImageInputFromURL is a wrapper function creating a single image input by URL.
func ImageInputFromURL(s string, m *interface{}) *Input {

	imd := &ImageData{}
	imd.AddFromURL(s)

	ind := &InputData{}
	ind.AddImage(imd)

	if m != nil {
		ind.AddMetadata(m)
	}

	return &Input{
		Data: ind,
	}
}

// ImageInputFromPath adds an image input based on a locally stored image.
func ImageInputFromPath(s string) (*Input, error) {

	out := &Input{}

	imd := &ImageData{}
	err := imd.AddFromBase64(s)
	if err != nil {
		return out, err
	}

	ind := &InputData{}
	ind.AddImage(imd)

	return &Input{
		Data: ind,
	}, nil
}

type InputService struct {
	Session *Session
	Page    int
	PerPage int
}

func (svc *InputService) GetPage() int {
	return svc.Page
}

func (svc *InputService) GetPerPage() int {
	return svc.PerPage
}

func (svc *InputService) GetSession() *Session {
	return svc.Session
}

func (svc *InputService) GetInputObject() *InputObject {
	return nil
}

func NewInputService(s *Session) *InputService {
	return &InputService{
		Session: s,
	}
}

type InputResponse struct {
	Status *ServiceStatus `json:"status,omitempty"`
	Inputs []*Input       `json:"inputs,omitempty"`
}

func (r *InputResponse) GetStatus() *ServiceStatus {
	return r.Status
}

//
//
// API
//
//

func (svc *InputService) WithPagination(page, perPage int) {
	svc.Page = page
	svc.PerPage = perPage
}

func (svc *InputService) ListAllInputs() (*InputResponse, error) {
	out := &InputResponse{}

	err := GetCall(svc, GetURI(svc, "inputs"), out)
	if err != nil {
		return out, err
	}

	return out, nil
}
