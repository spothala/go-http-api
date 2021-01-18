package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CreateRequest - Creates the http.Request object with necessary headers
func (httpW *HTTPWrapper) CreateRequest(method string, header map[string][]string, url string, body io.Reader) (*http.Request, error) {
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		httpReq.Header = header
	}
	if httpW.debug {
		fmt.Println("Requested URL: " + url)
	}
	return httpReq, nil
}

// ProcessBasicAuthRequest - Processes request with basic authorization & returns response
func (httpW *HTTPWrapper) ProcessBasicAuthRequest(method, username, password string, header map[string][]string,
	apiURL string, body io.Reader) (map[string][]string, []byte, int, error) {
	httpReq, err := httpW.CreateRequest(method, header, apiURL, body)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}
	httpReq.SetBasicAuth(username, password)
	if httpW.debug {
		fmt.Println(header)
		fmt.Println(httpReq.URL)
	}
	resp, err := httpW.client.Do(httpReq)
	if err != nil {
		return nil, nil, http.StatusNotFound, err
	}
	respBody, err := ReturnResponseBody(resp)
	return resp.Header, respBody, resp.StatusCode, err
}

// ProcessRequest - Processes basic request & returns response
func (httpW *HTTPWrapper) ProcessRequest(method string, header map[string][]string,
	apiURL string, body io.Reader) (map[string][]string, []byte, int, error) {
	httpReq, err := httpW.CreateRequest(method, header, apiURL, body)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}
	if httpW.debug {
		fmt.Println(header)
		fmt.Println(httpReq.URL)
	}
	resp, err := httpW.client.Do(httpReq)
	if err != nil {
		return nil, nil, http.StatusNotFound, err
	}
	respBody, err := ReturnResponseBody(resp)
	return resp.Header, respBody, resp.StatusCode, err
}

// ReturnResponseBody - Reads the response body & returns in byte format
func ReturnResponseBody(httpResponse *http.Response) ([]byte, error) {
	//defer body.Close()
	if httpResponse.ContentLength != 0 {
		contents, err := ioutil.ReadAll(httpResponse.Body)
		return contents, err
	}
	return []byte(""), nil
}
