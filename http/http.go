package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetRequestJsonParam(r *http.Request) (map[string]interface{}, error) {
	data, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		return nil, err
	}

	m, err := DecodeJson(data)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func EncodeJson(data interface{}) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DecodeReader(r io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}

func DecodeJson(data []byte) (map[string]interface{}, error) {
	var m map[string]interface{}

	err := json.Unmarshal(data, &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func RequestHttpJson(method, url string, body io.Reader) ([]byte, int, error) {
	headers := make(map[string][]string)
	headers["Content-Type"] = []string{"application/json"}
	return RequestHttp(method, "", url, body, headers)
}

func RequestHttp(method, baseUrl, path string, body io.Reader, headers map[string][]string) ([]byte, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, baseUrl+path, body)
	if err != nil {
		return nil, 408, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	dataBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	return dataBody, resp.StatusCode, nil
}

func Response(data []byte, statusCode int, resp http.ResponseWriter) {
	resp.WriteHeader(statusCode)
	fmt.Fprintf(resp, string(data))
}
