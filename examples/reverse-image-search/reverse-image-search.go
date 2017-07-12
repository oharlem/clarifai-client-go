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

	q := cl.NewAndSearchQuery()

	// Search by image URL.
	i := cl.NewImageFromURL("https://samples.clarifai.com/metro-north.jpg")

	// Search by a local file.
	//i, err := cl.NewImageFromFile("../Dave_Gahan_New_York_2015-10-22.jpg")
	//if err != nil {
	//	panic(err)
	//}

	q.WithImage(i)

	resp, err := sess.Search(q).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)

}
