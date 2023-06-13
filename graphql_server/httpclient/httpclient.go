package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func testHttpPost() (string, error) {
	posturl := ""

	// JSON body
	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 1
	}`)

	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPost, posturl, bytes.NewBuffer(body))
	if err != nil {
		return "", nil
	}

	resp := req.Response

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), err
}

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	// resp.StatusCode
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func HttpPost(url string, contentType string, body []byte) (string, error) {
	// application/json

	resp, err := http.Post(url, contentType, bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
