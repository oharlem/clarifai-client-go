# Clarifai Client for Go

[![Clarifai Client Go](https://goreportcard.com/badge/github.com/mpmlj/clarifai-client-go)](https://goreportcard.com/report/github.com/mpmlj/clarifai-client-go) [![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/mpmlj/clarifai-client-go/LICENSE.txt) [![GoDoc](https://godoc.org/github.com/mpmlj/clarifai-client-go?status.svg)](https://godoc.org/github.com/mpmlj/clarifai-client-go)

clarifai-client-go is an unofficial Go client for [Clarifai](https://www.clarifai.com/), an amazing, powerful AI image and video recognition service. 

## Clairfai API support

Only version 2.0 of Clarifai API is supported!

## Functionality
Current version of the client supports Predict calls only with ability to:
 - set a specific model
 - add an image input from URL
 - add an image input from a local file referenced by a path

## Basic Use

Install the client:

```
go get -u github.com/mpmlj/clarifai-client-go
```

Start using as:

```Go
package main

import (
        "fmt"

        client "github.com/mpmlj/clarifai-client-go"
)

func main() {
        sess := client.NewSession("app-client-id", "app-client-secret")
        err := sess.Connect()
        if err != nil {
        	panic(err.Error())
        }
        
        svc := client.NewPredictService(sess)
	    svc.SetModel(client.PublicModelFood)
	    
	    _ = svc.AddInput(client.ImageInputFromURL("https://samples.clarifai.com/food.jpg"))
        
        resp, err := svc.Call()
        if err != nil {
        	panic(err.Error())
        }
        
        fmt.Printf("%+v", resp.Outputs)
        	
}
```

## Demo
For a more complete demo check out a Feely application from [https://github.com/mpmlj/feely](https://github.com/mpmlj/feely). 

## Roadmap

- Token refresh on expiry
- Pagination support
- Image type validation as per [https://developer-preview.clarifai.com/guide/#supported-types](https://developer-preview.clarifai.com/guide/#supported-types)
- Search
- Training
- Advanced input management
- Advanced model management
- Advanced search functions
 