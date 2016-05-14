package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func NewClient() (client *http.Client) {
	// Creating HTTP Client with SSL support - Its Secure but we'll skip cert verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return
}

func CreateRequest(method string, header map[string][]string, url string, body io.Reader) (client *http.Client, httpReq *http.Request) {
	client = NewClient()
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("Failed to Prepare HTTP Request")
	}
	if header != nil {
		httpReq.Header = header
	}
	//	fmt.Println("Requested URL: " + url)
	return
}

func ProcessFormRequest(method string, header map[string][]string, Apiurl string, data url.Values) (response []byte) {
	httpReq, err := http.NewRequest(method, Apiurl, strings.NewReader(data.Encode()))
	if header != nil {
		httpReq.Header = header
	}
	if err != nil {
		fmt.Println("Failed to Prepare JsonRequest")
	}
	resp, err := NewClient().Do(httpReq)
	if err != nil {
		fmt.Println(err)
	}
	checkHttpResponseStatus(resp)
	return ReturnResponseBody(resp)
}

func ProcessRequest(method string, header map[string][]string, Apiurl string, body io.Reader) (response []byte, statusCode int) {
	client, httpReq := CreateRequest(method, header, Apiurl, body)
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
	}
	checkHttpResponseStatus(resp)
	return ReturnResponseBody(resp), resp.StatusCode
}

func checkHttpResponseStatus(httpResponse *http.Response) {
	// Following Check is for all Success Codes 2XX
	if !strings.HasPrefix(strconv.Itoa(httpResponse.StatusCode), "2") {
		PrintHttpResponseBody(httpResponse)
	}
}

func PrintHttpResponseBody(httpResponse *http.Response) {
	//defer httpResponse.Body.Close()
	contents, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func ReturnResponseBody(httpResponse *http.Response) (response []byte) {
	//defer body.Close()
	if httpResponse.ContentLength != 0 {
		contents, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return contents
	}
	return []byte("")
}

func RespondError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), http.StatusNotFound)
	return
}
