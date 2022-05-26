package afcclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Default URL
const HostURL string = "http://localhost:8000"
const Token string = "137f049286bb6330a30cd2ca4926ba072857a825eba0b397377e875c28efbfb9ffde16c87188d9b9355352b14bcd9d1f"

type Client struct {
	HostURL    string
	Token      string
	HTTPClient *http.Client
}

func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
		Token:      Token,
	}
	if host != nil {
		c.HostURL = *host
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
