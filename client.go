package afcclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Default URL
const HostURL string = "http://localhost:8000"
const Token string = "abcd"

//const host string = "http://localhost:8000"
//const token string = "abcd"
//HostURL = *host
//Token = *token
//

type Client struct {
	HTTPClient *http.Client
	HostURL    string
	Token      string
}

func NewClient(host *string, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
	c.HostURL = *host
	c.Token = *token
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
