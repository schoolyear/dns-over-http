package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func resolve(name string, queryType string) (data []byte, status int, err error) {
	// build URL
	query := url.Values{}
	query.Set("name", name)
	query.Set("type", queryType)
	reqUrl := url.URL{
		Scheme:   "https",
		Host:     "cloudflare-dns.com",
		Path:     "/dns-query",
		RawQuery: query.Encode(),
	}

	// build request
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Accept", "application/dns-json")

	// perform request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("failed to close resolve body", err)
		}
	}()

	// read response
	data, err = ioutil.ReadAll(res.Body)

	return data, res.StatusCode, err
}
