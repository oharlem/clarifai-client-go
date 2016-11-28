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

	q := cl.NewSearchQuery(cl.SearchQueryTypeAnd)
	q.WithUserConcept("album")  // inputs
	//q.WithoutAPIConcept("dog") // outputs
	q.SetPagination(1, 5)

	resp, err := sess.Search(q)
	if err != nil {
		panic(err)
	}

	cl.PP(resp)

}
