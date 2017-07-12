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

	// Map of desired concept changes as "concept id" -> bool state.
	concepts := map[string]bool{
		"album":    false,
		"vacation": true,
	}

	resp, err := sess.UpdateInputConcepts("travel-1", concepts).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
