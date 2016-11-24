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

	// Search By Custom Metadata.
	svc := cl.NewSearchService(sess)

	// Sample metadata.
	m := map[string]int{
		"version": 1,
	}

	resp, err := svc.SearchByCustomMetadata(m)
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
