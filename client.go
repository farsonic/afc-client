package afcclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Default URL
const HostURL string = "http://localhost:8000"
const Token string = "1234"

//

type Client struct {
	HostURL    string
	Token      string
	HTTPClient *http.Client
}

func NewClient(host *string) (token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
		Token:      Token,
	}
	if host != nil {
		c.HostURL = *host
	}
	if token != nil {
		c.Token = *token
	}
	return &c, nil
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
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
