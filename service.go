package clarifai

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Servicer interface {
	GetSession() *Session
	GetInputObject() *InputObject
	GetPage() int
	GetPerPage() int
}

type ServiceStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

// GetPayload prepares a service request payload.
func GetPayload(svc Servicer) (io.Reader, error) {

	var payload io.Reader
	body, err := json.Marshal(svc.GetInputObject())
	if err != nil {
		return payload, err
	}

	payload = bytes.NewReader(body)

	return payload, nil
}

func PrepPayload(i interface{}) (io.Reader, error) {

	var payload io.Reader
	body, err := json.Marshal(i)
	if err != nil {
		return payload, err
	}

	payload = bytes.NewReader(body)

	return payload, nil
}

// GetCall is a wrapper for HTTPCall for GET requests.
func GetCall(svc Servicer, endpoint string, resp interface{}) error {
	return HTTPCall(svc, http.MethodGet, endpoint, nil, resp)
}

// PostCall is a wrapper for HTTPCall for POST requests.
func PostCall(svc Servicer, endpoint string, payload io.Reader, resp interface{}) error {
	return HTTPCall(svc, http.MethodPost, endpoint, payload, resp)
}

// HTTPCall is a universal service caller with (re-)authentication and unmarshalling.
func HTTPCall(svc Servicer, method, endpoint string, payload io.Reader, resp interface{}) error {

	s := svc.GetSession()

	// Check for token expiration. If expired, re-authorize.
	if s.IsTokenExpired() {
		err := s.Connect()
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+s.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
//PP(string(body))
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}

	return nil
}

// GetURI generates a service endpoint.
func GetURI(svc Servicer, endpoint string) string {

	uri := svc.GetSession().GetAPIHost(endpoint)

	if svc.GetPage() > 0 || svc.GetPerPage() > 0 {
		v := url.Values{}
		v.Set("page", strconv.Itoa(svc.GetPage()))
		v.Add("per_page", strconv.Itoa(svc.GetPerPage()))
		return uri + "?" + v.Encode()
	}

	return uri
}
