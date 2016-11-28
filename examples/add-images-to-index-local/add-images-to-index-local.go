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

	// Create a new request.
	r := cl.NewRequest(sess)

	// Prepare a new image from URL.
	i, err := cl.NewImageFromFile("../Dave_Gahan_New_York_2015-10-22.jpg")
	if err != nil {
		panic(err)
	}

	// Add a concepts.
	i.AddConcepts("badn", "Dave Gahan", "Depeche Mode")

	// Allow adding duplicate images (off by default).
	i.AllowDuplicates()

	// Add custom metadata.
	m := map[string]string{
		"event_type": "show",
	}
	i.AddMetadata(m)

	// Add image to request.
	_ = r.AddImageInput(i)

	// Send request.
	resp, err := sess.AddImagesToIndex(r)
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
