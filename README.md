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

  
#### Input calls
- Add an image input from URL
- Add an image input from a local file
- Add image with concepts
- Add image with custom metadata
- Add image with crop
- Get a list of all inputs
- Get input by ID
- Get input status
- Get status of all inputs
- Input update adding concepts
- Input update deleting concepts
- Delete single input by ID
- Delete multiple inputs
- Delete all inputs


#### Models
- Create a model
- Get all models
- Get a model by id
- Get model output info
- Get all model versions
- Get model version by version ID
- Get all model inputs
- Get model inputs used to train a specific version
- Delete model
- Delete model version
- Delete all models
- Model training
- Add model concepts
- Delete model concepts
- Model search by name and/or type


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
Check directory ``/examples`` for fully-functional examples.
Please note, you need to set two env variables first:
- CLARIFAI_API_ID
- CLARIFAI_API_SECRET



## Roadmap

- Improve test coverage

 
## Support

- Go versions: 1.6, 1.7
- Clarifai API: 2.0