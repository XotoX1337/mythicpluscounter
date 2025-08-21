package raiderio

import (
	"encoding/json"
	"github.com/XotoX1337/mythicpluscounter/app"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type Client struct {
	client    *http.Client
	service   service
	Runs      *RunsService
	BaseUrl   *url.URL
	UserAgent string
	Token     string
}

type service struct {
	client *Client
}

type ListOptions struct {
	Page int `url:"page,omitempty"`
}

func NewClient() (*Client, error) {
	var err error
	c := &Client{client: &http.Client{}}
	c.service.client = c

	c.BaseUrl, err = url.Parse("https://raider.io/api/")
	c.UserAgent = app.Get().UserAgent()
	if err != nil {
		return nil, err
	}

	// convert pointer of service to the corresponding Service Struct
	c.Runs = (*RunsService)(&c.service)

	return c, nil
}

func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func (c *Client) NewRequest(method, url string) (*http.Request, error) {
	queryUrl, err := c.BaseUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	//req.Header.Set("PRIVATE-TOKEN", c.Token)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	} else {
		app.PrintInfo("GET %s", req.URL)
	}

	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decodeError := json.NewDecoder(resp.Body).Decode(v)
		if decodeError != nil {
			err = decodeError
		}
	}

	return err
}
