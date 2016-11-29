package clarifai

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiHost    = "https://api.clarifai.com"
	apiVersion = "v2"
)

var (
	userAgent string
)

func init() {
	userAgent = "clarifai-client-go/" + ClientVersion
}

type Session struct {
	clientID        string
	clientSecret    string
	accessToken     string
	tokenExpiration int
	host            string
}

type AuthResponse struct {
	AuthStatus  AuthStatus `json:"status"`
	AccessToken string     `json:"access_token"`
	ExpiresIn   int        `json:"expires_in"`
	Scope       string     `json:"scope"`
}

type AuthStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

// ServiceStatus is a universal status info object.
type ServiceStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Details     string `json:"details,omitempty"` // optional
}

// NewSession returns a default session object.
func Connect(clientID, clientSecret string) (*Session, error) {
	sess := NewSession(clientID, clientSecret)

	err := sess.Connect()
	if err != nil {
		return sess, err
	}

	return sess, nil
}

func NewSession(clientID, clientSecret string) *Session {
	return &Session{
		clientID:     clientID,
		clientSecret: clientSecret,
		host:         apiHost,
	}
}

// Connect contacts Clarifai API, tries to authenticate and returns access data on success.
func (s *Session) Connect() error {

	form := url.Values{}
	form.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, s.buildURI("token"), strings.NewReader(form.Encode()))
	req.SetBasicAuth(s.clientID, s.clientSecret)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

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

	var respObj AuthResponse
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return err
	}

	err = authResponseValidation(&respObj)
	if err != nil {
		return err
	}

	s.accessToken = respObj.AccessToken
	s.setTokenExpiration(respObj.ExpiresIn)
	return nil
}

// setTokenExpiration calculates session expiration time based on the current token.
func (s *Session) setTokenExpiration(exp int) (int, int) {
	startingTime := time.Now().Second()
	s.tokenExpiration = startingTime + exp

	return startingTime, s.tokenExpiration
}

// HTTPCall is a universal service caller with (re-)authentication and unmarshalling.
func (s *Session) HTTPCall(method, path string, payload interface{}) (*Response, error) {

	var resp *Response
	var err error
	var p io.Reader

	// Check for token expiration. If expired, re-authorize.
	if s.isTokenExpired() {
		err = s.Connect()
		if err != nil {
			return resp, err
		}
	}

	if payload != nil {
		p, err = prepPayload(payload)
		if err != nil {
			return resp, err
		}
	}
	req, err := http.NewRequest(method, s.buildURI(path), p)
	if err != nil {
		return resp, err
	}
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// buildURI constructs a full endpoint URI based of request path, API host and current API version.
func (s *Session) buildURI(endpoint string) string {
	return s.host + "/" + apiVersion + "/" + endpoint
}

// isTokenExpired checks if current authentication token has expired.
func (s *Session) isTokenExpired() bool {
	if s.tokenExpiration <= time.Now().Second() {
		return true
	}

	return false
}

// authResponseValidation validates response of the connection call.
func authResponseValidation(r *AuthResponse) error {

	if r.AccessToken == "" {
		return ErrNoAuthenticationToken
	}

	return nil
}

// prepPayload prepares any payload for the HTTP call.
func prepPayload(i interface{}) (io.Reader, error) {

	var payload io.Reader
	body, err := json.Marshal(i)
	if err != nil {
		return payload, err
	}

	return bytes.NewReader(body), nil
}
