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

	sess := cl.NewSession(viper.GetString("clarifai_api.client_id"), viper.GetString("clarifai_api.client_secret"))

	err = sess.Connect()
	if err != nil {
		panic(err)
	}

	// Get predictions for images.
	svc := cl.NewPredictService(sess)

	// By default, general model is used with ID "aaa03c23b3724a16a56b629203edc62c" as per
	// https://developer-preview.clarifai.com/guide/predict#predict
	_ = svc.AddInput(cl.ImageInputFromURL("https://samples.clarifai.com/metro-north.jpg", nil))
	_ = svc.AddInput(cl.ImageInputFromURL("https://samples.clarifai.com/puppy.jpeg", nil))
	_ = svc.AddInput(cl.ImageInputFromURL("https://samples.clarifai.com/food.jpg", nil))

	// ... but you can also set your own model.
	//svc.SetModel(cl.PublicModelFood)
	//_ = svc.AddInput(cl.ImageInputFromURL("https://samples.clarifai.com/food.jpg"))

	// NOTE. Currently API does not accept a mix of URL and base64 - based images!
	//i, err := cl.ImageInputFromPath("test_image.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//_ = svc1.AddInput(i)

	resp, err := svc.GetPredictions()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
