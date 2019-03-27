package httpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Header key value pairs.
type Header struct {
	Key, Value string
}

// HttpRequest contains the details of an http request.
type HttpRequest struct {
	verb, url string
	header    []Header
	Msg       interface{}
	Status
}

// GET sets the appropriate header HTTP verb and keypairs to connect by the GET
// method.
func (r *HttpRequest) GET(s string) {
	r.verb = "GET"
	r.url = s
	r.header = append(r.header, accAppJson())
}

// Status contains the error response code and description.
type Status struct {
	Code    int
	Message string
}

// String method returns a string representation of the contained data.
func (s Status) String() string {
	return strconv.Itoa(s.Code) + " " + s.Message
}

// getRespStatus saves the response status in the HttpRequest struct.
func getStatus(req HttpRequest, resp *http.Response) HttpRequest {

	req.Code = resp.StatusCode
	req.Message = http.StatusText(resp.StatusCode)

	return req
}

// sendRequest compiles and sends an http request.
func sendRequest(req HttpRequest, body []byte) (*http.Response, error) {

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

// respDecodeJSON reads and stores an http responce body.
func respDecodeJSON(req HttpRequest, resp *http.Response) (HttpRequest, error) {

	var msg json.RawMessage

	err := json.NewDecoder(resp.Body).Decode(&msg)
	if err != nil {
		return req, fmt.Errorf("Decode: %v", err)
	}
	req.Msg = msg
	return req, nil
}

// Request sends and recieves an HTTP requests, storing the responce
// status.
func Request(req HttpRequest, body []byte) (HttpRequest, error) {

	response, err := sendRequest(req, body)
	if err != nil {
		return req, fmt.Errorf("sendRequest: %v", err)
	}

	req = getStatus(req, response)

	req, err = respDecodeJSON(req, response)
	if err != nil {
		return req, fmt.Errorf("respDecodeJSON: %v", err)
	}

	return req, nil
}
