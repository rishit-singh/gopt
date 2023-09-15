package util;

import (
	"strings";
	"net/http";
);

type HttpRequest struct {
	Request *http.Request;

	Headers map[string]string; 	
}

func NewHttpRequest(method string, endpoint string, headers map[string]string, body string) (*HttpRequest, error) {
	request := new(HttpRequest);

	var err error;

	request.Request, err = http.NewRequest(method, endpoint, strings.NewReader(body));

	for key, value := range headers {
		request.Request.Header.Add(key, value);
	}

	return request, err;
}