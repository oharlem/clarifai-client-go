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

	// Get a list of all inputs.
	svc := cl.NewInputService(sess)

	svc.WithPagination(1, 5) // Start from page 1 with 5 items per page.
	resp, err := svc.ListAllInputs()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
