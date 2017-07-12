package main

import (
	"os"

	cl "github.com/mpmlj/clarifai-client-go"
)

func main() {

	var err error
	var sess *cl.Session
	if os.Getenv("CLARIFAI_API_KEY") != "" {
		sess = cl.NewApp(os.Getenv("CLARIFAI_API_KEY"))
	} else {
		sess, err = cl.Connect(os.Getenv("CLARIFAI_API_ID"), os.Getenv("CLARIFAI_API_SECRET"))
		if err != nil {
			panic(err)
		}
	}

	// Init an inputs object.
	data := cl.InitInputs()

	// Prepare a new image from URL.
	i := cl.NewImageFromURL("https://samples.clarifai.com/travel.jpg")

	// Add concepts.
	i.AddConcepts([]string{"album", "vacation"})

	// Add this image, even if its duplicate is found ("off" by default).
	i.AllowDuplicates()

	// Add custom metadata.
	m := map[string]string{
		"event_type": "vacation",
	}
	i.AddMetadata(m)

	// Add crop points.
	i.AddCrop(0.2, 0.4, 0.3, 0.6)

	// Add image to request.
	_ = data.AddInput(i, "travel-1")

	// ...or you can skip image id for it to be automatically generated for you:
	// _ = data.AddInput(i, "")

	// Send request.
	resp, err := sess.AddInputs(data).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
