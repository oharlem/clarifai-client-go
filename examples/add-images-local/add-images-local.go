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

	// Create a new request.
	data := cl.InitInputs()

	// ADD IMAGE 1
	// Prepare a new image from a local file.
	i, err := cl.NewImageFromFile("../Dave_Gahan_New_York_2015-10-22.jpg")
	if err != nil {
		panic(err)
	}

	// Add a concepts.
	i.AddConcepts([]string{"Dave Gahan", "Depeche Mode"})

	// Allow adding duplicate images (off by default).
	i.AllowDuplicates()

	// Add custom metadata.
	m := map[string]string{
		"event_type": "show",
	}
	i.AddMetadata(m)

	// Add image to request.
	err = data.AddInput(i, "dm-1")
	if err != nil {
		panic(err)
	}

	// ADD IMAGE 2
	// Prepare a new image from a local file.
	i2, err := cl.NewImageFromFile("../depeche_mode.jpg")
	if err != nil {
		panic(err)
	}
	i2.AddConcepts([]string{"band", "Depeche Mode"})
	err = data.AddInput(i2, "dm-2")
	if err != nil {
		panic(err)
	}

	// Send request.
	resp, err := sess.AddInputs(data).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
