# Clarifai Client for Go

[![Clarifai Client Go](https://goreportcard.com/badge/github.com/mpmlj/clarifai-client-go)](https://goreportcard.com/report/github.com/mpmlj/clarifai-client-go) [![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/mpmlj/clarifai-client-go/LICENSE.txt) [![Build Status](https://travis-ci.org/mpmlj/clarifai-client-go.svg)](https://travis-ci.org/mpmlj/clarifai-client-go) [![GoDoc](https://godoc.org/github.com/mpmlj/clarifai-client-go?status.svg)](https://godoc.org/github.com/mpmlj/clarifai-client-go)

clarifai-client-go is an unofficial Go client for [Clarifai](https://www.clarifai.com/), an amazing, powerful AI image and video recognition service. 


## Functionality

#### General 
- Token refresh on expiry
- Pagination support

#### Predict calls
- Get predictions 
- With a specific model
- Add an image input from URL
- Add an image input from a local file
  
#### Input calls
- Get a list of all inputs
 
#### Search
- Add images to a search index
- Search by predicted concepts
- Search by user supplied concept
- Reverse image search
- Search by custom metadata
- Mixed search by concepts and predictions 
 
 
## Installation

```
go get -u github.com/mpmlj/clarifai-client-go
```

## Examples
Check directory /examples for fully-functional examples.
_You need to update conf.yml with your Clarifai API credentials._


## Roadmap

- Training
- Advanced input management
- Advanced model management

 
## Support

- Go versions: 1.6, 1.7
- Clarifai API: 2.0