package clarifai

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiHost    = "https://api.clarifai.com"
	apiVersion = "v2"
	userAgent  = "clarifai-client-go/0.2"
)

type Session struct {
	ClientID        string
	ClientSecret    string
	AccessToken     string
	TokenExpiration int
	Host            string
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

// NewSession returns a default session object.
func NewSession(clientID, clientSecret string) *Session {
	return &Session{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Host:         apiHost,
	}
}

func (s *Session) SetHost(h string) {
	s.Host = h
}

// Connect contacts Clarifai API, tries to authenticate and returns access data on success.
func (s *Session) Connect() error {

	form := url.Values{}
	form.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, s.GetAPIHost("token"), strings.NewReader(form.Encode()))
	req.SetBasicAuth(s.ClientID, s.ClientSecret)
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

	err = AuthResponseValidation(&respObj)
	if err != nil {
		return err
	}

	s.setAccessToken(respObj.AccessToken)
	s.SetTokenExpiration(respObj.ExpiresIn)
	return nil
}

// setAccessToken sets access token to a current session.
func (s *Session) setAccessToken(token string) {
	s.AccessToken = token
}

// GetAccessToken gets access token of a current session.
func (s *Session) GetAccessToken() string {
	return s.AccessToken
}

// SetTokenExpiration calculates session expiration time based on the current token.
func (s *Session) SetTokenExpiration(exp int) (int, int) {
	startingTime := time.Now().Second()
	s.TokenExpiration = startingTime + exp

	return startingTime, s.TokenExpiration
}

// IsTokenExpired checks if current authentication token has expired.
func (s *Session) IsTokenExpired() bool {
	if s.TokenExpiration <= time.Now().Second() {
		return true
	}

	return false
}

// GetAPIHost constructs a full endpoint URI.
func (s *Session) GetAPIHost(endpoint string) string {
	return s.Host + "/" + apiVersion + "/" + endpoint
}

// AuthResponseValidation validates response of the connection call.
func AuthResponseValidation(r *AuthResponse) error {

	if r.AccessToken == "" {
		return ErrNoAuthenticationToken
	}

	return nil
}
