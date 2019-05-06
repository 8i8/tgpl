package github

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

// Header request key pairs.
type Header struct {
	Key, Value string
}

// Address contains the URL the request and the response.
type Address struct {
	HTTP, URL string
	header    []Header
	Status
}

// Status http response, text and code key pair.
type Status struct {
	Code    int
	Message string
}

// String returns a string that contains the key pair.
func (s Status) String() string {
	return strconv.Itoa(s.Code) + " " + s.Message
}

// sendRequest compiles and sends the appropriate predefined request.
func sendRequest(addr Address, json []byte) (*http.Response, error) {

	// Buffer json for request.
	buf := bytes.NewReader(json)

	// Get a new request object.
	req, err := http.NewRequest(addr.HTTP, addr.URL, buf)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %v", err)
	}

	// Load the key value pairs into the request.
	for _, h := range addr.header {
		req.Header.Set(h.Key, h.Value)
	}

	// Make the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient: %v", err)
	}
	return resp, err
}

// getStatus fills a Status struct from an http response status and verifies
// that response data against the expected, raising an error on discrepancies.
func getStatus(resp *http.Response) (Status, error) {

	var s Status
	s.Code = resp.StatusCode
	s.Message = http.StatusText(resp.StatusCode)

	// Check mode and verify correct response status, raise error if
	// required.
	if f&(cLIST|cREAD|cEDIT) > 0 &&
		s.Code != http.StatusOK {
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
	} else if f&cRAISE > 0 && s.Code != http.StatusCreated {
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
	} else if f&(cLOCK|cUNLOCK) > 0 && s.Code != http.StatusNoContent {
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
	}

	return s, nil
}

// makeRequest orchestrates an http request.
func makeRequest(c Config, json []byte) (Reply, error) {

	var reply Reply
	// Set the correct url for the request.
	c, addr, err := setURL(c)
	if err != nil {
		return reply, fmt.Errorf("setURL: %v", err)
	}

	// Compile an array of header key value pairs.
	addr.header, err = composeHeader(c)
	if err != nil {
		return reply, fmt.Errorf("composeHeader: %v", err)
	}

	if f&cVERBOSE > 0 {
		fmt.Printf("makeRequest: json %v\n", string(json))
	}

	resp, err := sendRequest(addr, json)
	if err != nil {
		return reply, fmt.Errorf("sendRequest: %v", err)
	}

	// Record the response return state.
	addr.Status, err = getStatus(resp)
	if err != nil {
		return reply, fmt.Errorf("getStatus: %v", err)
	}

	if f&cVERBOSE > 0 {
		fmt.Printf("getStatus: %v\n", addr.String())
	}

	// Decode the response if required.
	if f&(cLIST|cREAD) > 0 {
		reply, err = respDecode(c, resp)
		if err != nil {
			return reply, fmt.Errorf("respDecode: %v", err)
		}
	}
	resp.Body.Close()

	return reply, err
}
