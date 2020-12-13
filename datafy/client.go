package datafy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Token struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int32  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      Token
}

func NewClient(host, profileName string) (*Client, error) {
	var token Token

	profilePath := fmt.Sprintf("%v/.datafy/tokens/%v", os.Getenv("HOME"), profileName)
	fmt.Println(profilePath)

	profileValue, err := ioutil.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(profileValue, &token); err != nil {
		return nil, err
	}
	fmt.Println(token)

	c := Client{
		HostURL:    host,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      token,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token.AccessToken)
	req.Header.Set("accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
