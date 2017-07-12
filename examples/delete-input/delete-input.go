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

	// Delete a single input by its id.
	resp, err := sess.DeleteInput("music-2").Do()
	if err != nil {
		panic(err)
	}

	// Delete multiple inputs by their ids.
	//resp, err := sess.DeleteInputs([]string{"music-1", "music-2"}).Do()
	//if err != nil {
	//	panic(err)
	//}

	// Delete all inputs.
	//resp, err := sess.DeleteAllInputs().Do()
	//if err != nil {
	//	panic(err)
	//}

	cl.PP(resp)
}
