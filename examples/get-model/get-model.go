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

	//// Get a single model by id with minimum data.
	//resp, err := sess.GetModel("music-model-id-1").Do()
	//if err != nil {
	//	panic(err)
	//}

	//// Get a single model by id with output data.
	resp, err := sess.GetModelOutput("music-model-id-1").Do()
	if err != nil {
		panic(err)
	}

	//// Get a single model by id with versions data.
	//resp, err := sess.GetModelVersions("eab1fd01a5544225b32d5d2937e05041").Do()
	//if err != nil {
	//	panic(err)
	//}

	//// Get a single model version.
	//resp, err := sess.GetModelVersion("eab1fd01a5544225b32d5d2937e05041", "d88847bb75514fceaf74bf36606d1343").Do()
	//if err != nil {
	//	panic(err)
	//}

	// GetModelVersionInputs fetches inputs used to train a specific model version.
	//resp, err := sess.GetModelVersionInputs("eab1fd01a5544225b32d5d2937e05041", "d88847bb75514fceaf74bf36606d1343").Do()
	//if err != nil {
	//	panic(err)
	//}

	//// Get a single model inputs.
	//resp, err := sess.GetModelInputs("eab1fd01a5544225b32d5d2937e05041").Do()
	//if err != nil {
	//	panic(err)
	//}

	cl.PP(resp)
}
