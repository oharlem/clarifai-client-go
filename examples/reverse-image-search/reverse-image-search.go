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

	// Reverse Image Search.
	svc := cl.NewSearchService(sess)

	resp, err := svc.ReverseImageSearch(
		cl.ImageInputFromURL("https://samples.clarifai.com/metro-north.jpg", nil),
		cl.ImageInputFromURL("https://samples.clarifai.com/puppy.jpeg", nil),
	)
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
