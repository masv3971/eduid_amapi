package eduid_amapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/masv/eduid_amapi/amapi_types"
)

type Config struct {
	URL string `validate:"required"`
}

type Client struct {
	HTTPClient *http.Client
	url        string

	User    *userService
	Sampler *samplerService
}

// New create a new instance of eduid_amapi
func New(config *Config) *Client {
	c := &Client{
		url: config.URL,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}

	c.User = &userService{
		client:     c,
		subBaseURL: "/users/",
	}
	c.Sampler = &samplerService{
		client:     c,
		subBaseURL: "/sampler/",
	}

	return c
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	completeURL := u.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		payload := struct {
			Data interface{} `json:"data"`
		}{
			Data: body,
		}
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, completeURL.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	//req.Header.Set("Accept", acceptHeader)
	req.Header.Set("User-Agent", "eduid_amapi/0.0.1")

	return req, nil

}

func (c *Client) do(ctx context.Context, req *http.Request, reply interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		error := &amapi_types.AmAPIError{}
		if err := json.Unmarshal(buf.Bytes(), error); err != nil {
			return nil, err
		}
		return nil, error
	}

	return resp, nil

}

func checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return errors.New("invalid")
	}
	return errors.New("invalid request")
}

func (c *Client) call(ctx context.Context, method, path string, body, reply interface{}) (*http.Response, error) {
	request, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, request, reply)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
