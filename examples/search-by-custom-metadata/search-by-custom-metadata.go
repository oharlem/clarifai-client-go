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

	// Search By Custom Metadata.
	q := cl.NewSearchQuery(cl.SearchQueryTypeAnd)

	// Sample metadata.
	m := map[string]interface{}{
		"event_type": "wedding",
	}
	q.WithMetadata(m)

	resp, err := sess.Search(q)
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
