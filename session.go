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
	apiHost         = "https://api.clarifai.com"
	apiVersion      = "v2"
	timestampFormat = time.RFC3339
	userAgent       = "clarifai-client-go/0.1"
)

type Session struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func NewSession(clientID, clientSecret string) *Session {
	return &Session{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

// Connect contacts Clarifai API, tries to authenticate and returns access data on success.
func (c *Session) Connect() error {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, getAPIHost("token"), strings.NewReader(form.Encode()))
	req.SetBasicAuth(c.ClientID, c.ClientSecret)
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

	respObj := AuthResponse{}
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return err
	}

	c.setAccessToken(respObj.AccessToken)
	return nil
}

func (c *Session) setAccessToken(token string) {
	c.AccessToken = token
}

func (c *Session) GetAccessToken() string {
	return c.AccessToken
}

func getAPIHost(endpoint string) string {
	return apiHost + "/" + apiVersion + "/" + endpoint
}
