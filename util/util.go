package util;

import (
	"strings";
	"errors";
	"net/http";
	"encoding/json";
);

type any = interface{};

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

func InRange(val int, low int, high int) bool {
	return val >= low && val < high;
}

func CombineStrings(strs []string, delimiter string, start int, end int) (any, error) {
	size := len(strs);
	
	combined := "";

	if (!InRange(end, 0, size) && !InRange(start, 0, end + 1)) {
		return nil, errors.New("Index out of range");
	}
	
	
	for x := start; x < end; x++ {
		combined  += strs[x] + " ";
	}
		
	combined += strs[end];

	return combined, nil;
}

func JsonToMap(jsonStr string) (any, error) {	
	var jsonMap map[string]any;

	err := json.Unmarshal([]byte(jsonStr), &jsonMap);
		
	if (err != nil) {
		return nil, err;
	}
	
	return jsonMap, nil;
}	

