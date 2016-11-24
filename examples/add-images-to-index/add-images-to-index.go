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

	// Add images to search index.
	svc := cl.NewSearchService(sess)

	_ = svc.AddInput(cl.ImageInputFromURL("https://samples.clarifai.com/food.jpeg", nil))
	_ = svc.AddInput(cl.ImageInputFromURL("https://samples.clarifai.com/wedding.jpeg", nil))
	resp, err := svc.AddImagesToIndex() // response can be ignored
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
