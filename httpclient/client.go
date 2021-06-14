package httpclient

import (
	"fmt"
	"github.com/franela/goreq"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseUri string
}

func NewClient(baseUri string) *Client {
	return &Client{
		baseUri: baseUri,
	}
}

type RequestBuilder struct {
	request goreq.Request
	query   *url.Values
}

func (r *RequestBuilder) Param(param string, value string) *RequestBuilder {
	if "" != value {
		r.query.Set(param, value)
	}
	return r
}

func (r *RequestBuilder) Build() goreq.Request {
	r.request.QueryString = r.query
	return r.request
}

func (c *Client) GetRequest(paths ...string) *RequestBuilder {
	return &RequestBuilder{
		request: goreq.Request{
			Uri:             c.url(paths...),
			Accept:          "application/json",
			Insecure:        true,
			OnBeforeRequest: c.OnBeforeRequest,
		},
		query: &url.Values{},
	}
}

func (c *Client) ExternalGetRequest(uri string) *RequestBuilder {
	return &RequestBuilder{
		request: goreq.Request{
			Uri:             uri,
			Accept:          "application/json",
			Insecure:        true,
			OnBeforeRequest: c.OnBeforeRequest,
		},
		query: &url.Values{},
	}
}

func (c *Client) OnBeforeRequest(goreq *goreq.Request, httpreq *http.Request) {
	httpreq.Close = true
}

func (c *Client) PostRequest(body interface{}, paths ...string) *RequestBuilder {
	return &RequestBuilder{
		request: goreq.Request{
			Method:          "POST",
			ContentType:     "application/json",
			Accept:          "application/json",
			Uri:             c.url(paths...),
			Body:            body,
			Insecure:        true,
			OnBeforeRequest: c.OnBeforeRequest,
		},
		query: &url.Values{},
	}
}

func (c *Client) UploadRequest(body interface{}, contentType string, paths ...string) *RequestBuilder {
	return &RequestBuilder{
		request: goreq.Request{
			Method:          "POST",
			ContentType:     contentType,
			Accept:          "application/json",
			Uri:             c.url(paths...),
			Body:            body,
			Insecure:        true,
			OnBeforeRequest: c.OnBeforeRequest,
		},
		query: &url.Values{},
	}
}

func (c *Client) PutRequest(body interface{}, paths ...string) *RequestBuilder {
	return &RequestBuilder{
		request: goreq.Request{
			Method:          "PUT",
			ContentType:     "application/json",
			Accept:          "application/json",
			Uri:             c.url(paths...),
			Body:            body,
			Insecure:        true,
			OnBeforeRequest: c.OnBeforeRequest,
		},
		query: &url.Values{},
	}
}

func (c *Client) DeleteRequest(paths ...string) *RequestBuilder {
	return &RequestBuilder{
		request: goreq.Request{
			Method:          "DELETE",
			ContentType:     "application/json",
			Accept:          "application/json",
			Uri:             c.url(paths...),
			Insecure:        true,
			OnBeforeRequest: c.OnBeforeRequest,
		},
		query: &url.Values{},
	}
}

func (c *Client) DoExpectCodes(request *RequestBuilder, statuses ...int) (*goreq.Response, error) {
	return c.doExpectCodes(request.Build(), statuses...)
}

func (c *Client) doExpectCodes(request goreq.Request, statuses ...int) (*goreq.Response, error) {
	response, err := request.Do()
	if nil != err {
		return nil, fmt.Errorf("error executing request: %s", err)
	}
	for _, status := range statuses {
		if status == response.StatusCode {
			return response, nil
		}
	}
	defer response.Body.Close()
	body, _ := response.Body.ToString()
	return nil, fmt.Errorf("server returned code %d for %s %v -- response: %s",
		response.StatusCode, request.Method, request.Uri, body)
}

func (c *Client) DoExpectOk(request *RequestBuilder) (*goreq.Response, error) {
	return c.doExpectCodes(request.Build(), http.StatusOK)
}

func (c *Client) Unmarshal(response *goreq.Response, target interface{}) {
	defer response.Body.Close()
	if err := response.Body.FromJsonTo(target); err != nil {
		logrus.Fatalf("Error unmarshalling response: %s", err)
	}
}

func (c *Client) Execute(request *RequestBuilder, target interface{}) error {
	response, err := c.DoExpectOk(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err := response.Body.FromJsonTo(target); err != nil {
		return fmt.Errorf("error unmarshalling response: %s", err)
	}
	return nil
}

func (c *Client) url(paths ...string) string {
	fullPaths := append([]string{c.baseUri}, paths...)
	return strings.Join(fullPaths, "/")
}

func (c *Client) addIfNotBlank(query *url.Values, name string, value string) {
	if "" != value {
		query.Add(name, value)
	}
}

func (c *Client) addIfNotEmpty(query *url.Values, name string, value *[]string) {
	if nil == value || len(*value) == 0 {
		return
	}
	query.Add(name, strings.Join(*value, ","))
}
