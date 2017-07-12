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

	//resp, err := sess.DeleteModel("ee471ed8ee8849d8a15fa24f46d8814c").Do()
	//if err != nil {
	//	panic(err)
	//}

	//resp, err := sess.DeleteModelVersion("eab1fd01a5544225b32d5d2937e05041", "d88847bb75514fceaf74bf36606d1343").Do()
	//if err != nil {
	//	panic(err)
	//}

	resp, err := sess.DeleteAllModels().Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
