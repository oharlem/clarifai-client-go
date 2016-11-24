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

	// Search By Concept AND Predicted Concept.
	svc8 := cl.NewSearchService(sess)

	q := cl.NewQuery()

	q.AddOutputConcepts([]string{"album", "test"})
	q.AddInputConcepts([]string{"album"})

	resp8, err := svc8.Call(q)
	if err != nil {
		panic(err)
	}

	cl.PP(resp8)
}
