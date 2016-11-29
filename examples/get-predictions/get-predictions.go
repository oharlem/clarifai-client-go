package main

import (
	cl "github.com/mpmlj/clarifai-client-go"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("..")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	sess, err := cl.Connect(viper.GetString("clarifai_api.client_id"), viper.GetString("clarifai_api.client_secret"))
	if err != nil {
		panic(err)
	}

	r := cl.NewRequest(sess)

	// Option A. Adding an image from URL.
	_ = r.AddImageInput(cl.NewImageFromURL("https://samples.clarifai.com/metro-north.jpg"))

	// Option B. Adding an image from a local file.
	// NOTE. Currently API does not accept a mix of URL and base64 - based images!
	//i, err := cl.NewImageFromFile("../Dave_Gahan_New_York_2015-10-22.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//_ = r.AddInput(i)

	// By default, general model is used with ID "aaa03c23b3724a16a56b629203edc62c" as per
	// https://developer-preview.clarifai.com/guide/predict#predict, but you can also set your own model.
	//r.SetModel(cl.PublicModelTravel)

	resp, err := sess.GetPredictions(r)
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
