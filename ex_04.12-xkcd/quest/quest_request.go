package quest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var count uint

// respDecode reads and stores an http response body.
func respDecode(req *HttpQuest, resp *http.Response) error {

	var msg json.RawMessage

	err := json.NewDecoder(resp.Body).Decode(&msg)
	if err != nil {
		return fmt.Errorf("Decode: %v", err)
	}
	req.Msg = msg
	return nil
}

// getRespStatus saves the response status in the HttpQuest struct.
func getStatus(req *HttpQuest, resp *http.Response) *HttpQuest {

	req.Code = resp.StatusCode
	req.Message = http.StatusText(resp.StatusCode)

	if VERBOSE {
		count++
		i := count % 6
		switch i {
		case 0:
			fmt.Printf("\rquest: http: %s .  ", req.Status())
		case 1:
			fmt.Printf("\rquest: http: %s  . ", req.Status())
		case 2:
			fmt.Printf("\rquest: http: %s   .", req.Status())
		case 3:
			fmt.Printf("\rquest: http: %s   .", req.Status())
		case 4:
			fmt.Printf("\rquest: http: %s  . ", req.Status())
		case 5:
			fmt.Printf("\rquest: http: %s .  ", req.Status())
		}
	}

	return req
}

// sendRequest compiles and sends an http request.
func sendRequest(req HttpQuest, body []byte) (*http.Response, error) {

	// Buffer jsonl
	buf := bytes.NewReader(body)

	// Retrieve a new reaquest object.
	request, err := http.NewRequest(req.verb, req.url, buf)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %v", err)
	}

	// Load the key value pairs into the request.
	for _, h := range req.header {
		request.Header.Set(h.Key, h.Value)
	}

	// Make the request.
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient.Do: %v", err)
	}

	return response, err
}

// request sends and receives an HTTP requests, storing the response
// status.
func request(req HttpQuest, body []byte) (HttpQuest, error) {

	response, err := sendRequest(req, body)
	if err != nil {
		return req, fmt.Errorf("sendRequest: %v", err)
	}

	getStatus(&req, response)

	err = respDecode(&req, response)
	if err != nil && req.Code != 404 {
		return req, fmt.Errorf("respDecode: %v", err)
	}

	return req, nil
}
