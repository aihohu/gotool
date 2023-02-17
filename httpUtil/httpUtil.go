package httpUtil

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Header struct {
	Key   string
	Value string
}

var (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
)

// HttpGet GET 请求
func HttpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("write from conn failed, err:%v\n\n", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if len(body) == 0 {
		return nil
	}

	//log.Println("Get: ", len(body), string(body))
	return body
}

// HttpPost Post 请求
func HttpPost(urlStr string, data []byte) []byte {

	resp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(data))
	if err != nil {
		//log.Fatal(err)
		log.Printf("write from conn failed, err:%v\n\n", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Println("Post err : ", resp)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if len(body) == 0 {
		return nil
	}

	//log.Println("Post: ", len(body), string(body))
	return body
}

// HttpPut PUT 请求
func HttpPut(urlStr string, data []byte) []byte {

	req, _ := http.NewRequest("PUT", urlStr, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("write from conn failed, err:%v\n\n", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if len(body) == 0 {
		return nil
	}

	//log.Println("put: ", len(body), string(body))
	return body
}

func HttpRequest(method, url string) []byte {
	return HttpRequestHeader(method, url, nil, nil)
}

func HttpRequestBody(method, url string, body io.Reader) []byte {
	return HttpRequestHeader(method, url, body, nil)
}

func HttpRequestHeader(method, url string, body io.Reader, headers []Header) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("HttpRequest failed, err:%v\n\n", err)
		return nil
	}

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("HttpRequest failed, err:%v\n\n", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if len(respBody) == 0 {
		return nil
	}

	return respBody
}
