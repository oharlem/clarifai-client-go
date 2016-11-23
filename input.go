package clarifai

import (
	"io/ioutil"
	"os"
)

type InputsCollection struct {
	Inputs []*Input `json:"inputs"`
}

type Input struct {
	Data InputData `json:"data,omitempty"`
	ID   string    `json:"id,omitempty"`
}

type InputData struct {
	Concepts []*InputConcept `json:"concepts,omitempty"`
	Image    *ImageData      `json:"image,omitempty"`
	Metadata *Metadata       `json:"metadata,omitempty"`
}

func (id *InputData) AddImage(i ImageData) {
	id.Image = &i
}

type ImageData struct {
	Base64 string    `json:"base64,omitempty"`
	URL    string    `json:"url,omitempty"`
	Crop   []float32 `json:"crop,omitempty"`
}

type Metadata struct {
	Key  string `json:"key"`
	List []int  `json:"list"`
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

	i.Base64 = EncBytesToBase64(data)

	return nil
}

// ImageInputFromURL, a wrapper function, creates a single image input by URL.
func ImageInputFromURL(s string) *Input {

	imd := ImageData{}
	imd.AddFromURL(s)

	ind := InputData{}
	ind.AddImage(imd)

	return &Input{
		Data: ind,
	}
}

// ImageInputFromPath adds an image input based on a locally stored image.
func ImageInputFromPath(s string) (*Input, error) {

	out := &Input{}

	imd := ImageData{}
	err := imd.AddFromBase64(s)
	if err != nil {
		return out, err
	}

	ind := InputData{}
	ind.AddImage(imd)

	return &Input{
		Data: ind,
	}, nil
}
