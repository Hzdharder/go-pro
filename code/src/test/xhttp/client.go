package xhttp

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	HttpClient  *http.Client
	Transport   *http.Transport
	Header      http.Header
	url         string
	method      string
	requestType RequestType
	FormString  string
	JsonByte    []byte
	ContentType string
}

func NewClient() (client *Client) {
	return &Client{
		HttpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				DisableKeepAlives: true,
				Proxy:             http.ProxyFromEnvironment,
			},
		},
		Header: make(http.Header),
		requestType: TypeJSON,
	}
}

func (c *Client) SetTLSConfig(tlsCfg *tls.Config) (client *Client) {
	c.Transport = &http.Transport{TLSClientConfig: tlsCfg, DisableKeepAlives: true, Proxy: http.ProxyFromEnvironment}
	return c
}

func (c *Client) Type(typeStr RequestType) (client *Client) {
	if _, ok := types[typeStr]; ok {
		c.requestType = typeStr
	}
	return c
}

func (c *Client) Post(url string) (client *Client) {
	c.method = POST
	c.url = url
	return c
}

func (c *Client) Get(url string) (client *Client) {
	c.method = GET
	c.url = url
	return c
}

func (c *Client) SendBodyMap(bm map[string]interface{}) (client *Client) {
	if bm == nil {
		return c
	}
	switch c.requestType {
	case TypeJSON:
		bs, err := json.Marshal(bm)
		if err != nil {
			return c
		}
		c.JsonByte = bs
	case TypeXML, TypeForm:
		c.FormString = FormatURLParam(bm)
	}
	return c
}

func (c *Client) SendString(encodeStr string) (client *Client) {
	switch c.requestType {
	case TypeJSON:
		c.JsonByte = []byte(encodeStr)
	case TypeXML, TypeForm:
		c.FormString = encodeStr
	}
	return c
}

func (c *Client) EndBytes() (res *http.Response, bs []byte, err error) {
	var body io.Reader

	reqFunc := func() (err error) {
		switch c.method {
		case GET:
			switch c.requestType {
			case TypeJSON:
				c.ContentType = types[TypeJSON]
			case TypeForm:
				c.ContentType = types[TypeForm]
			case TypeXML:
				c.ContentType = types[TypeXML]
			default:
				return errors.New("Request Type Error ")
			}
		case POST:
			switch c.requestType {
			case TypeJSON:
				if c.JsonByte != nil {
					body = strings.NewReader(string(c.JsonByte))
				}
				c.ContentType = types[TypeJSON]
			case TypeForm:
				body = strings.NewReader(c.FormString)
				c.ContentType = types[TypeForm]
			case TypeXML:
				body = strings.NewReader(c.FormString)
				c.ContentType = types[TypeXML]
			default:
				return errors.New("Request Type Error ")
			}
		default:
			return errors.New("Only support GET and POST ")
		}

		req, err := http.NewRequest(c.method, c.url, body)
		if err != nil {
			return err
		}
		req.Header = c.Header
		req.Header.Set("Content-Type", c.ContentType)
		if c.Transport != nil {
			c.HttpClient.Transport = c.Transport
		}

		res, err = c.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		bs, err = ioutil.ReadAll(io.LimitReader(res.Body, int64(5<<20)))
		if err != nil {
			return err
		}
		return nil
	}
	if err = reqFunc(); err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}

func FormatURLParam(body map[string]interface{}) (urlParam string) {
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range body {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v, ok := body[k].(string)
		if !ok {
			v = convertToString(body[k])
		}
		if v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}