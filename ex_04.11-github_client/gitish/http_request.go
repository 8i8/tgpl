package gitish

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Header contains header request key pairs.
type Header struct {
	Key, Value string
}

// Status contains an http responce status, both the text and the code.
type Status struct {
	Code    int
	Message string
}

// String prints the details of an HTTP status struct.
func (s Status) String() string {
	return strconv.Itoa(s.Code) + " " + s.Message
}

// Address contains the request tupe and URL of a request.
type Address struct {
	HTTP, URL string
	header    []Header
	Status
}

// sendRequest compiles and sends the appropriate predefined request.
func sendRequest(conf Config, addr Address, json io.Reader) (*http.Response, error) {

	// Get a new request object.
	req, err := http.NewRequest(addr.HTTP, addr.URL, json)
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

// getStatus fills a Status struct with an http responce status data and
// verifies that responce data agains the expected responce, raising an error
// when required.
func getStatus(resp *http.Response) (Status, error) {

	var s Status
	s.Code = resp.StatusCode
	s.Message = http.StatusText(resp.StatusCode)

	if (f&cLIST > 0 || f&cREAD > 0 || f&cEDIT > 0) &&
		s.Code != http.StatusOK {
		// If data recieved.
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
		// If a record was created.
	} else if f&cRAISE > 0 && s.Code != http.StatusCreated {
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
	} else if (f&cLOCK > 0 || f&cUNLOCK > 0) && s.Code != http.StatusNoContent {
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
	}

	return s, nil
}

// makeRequest orchestrates an http request.
// TODO is the header being set in the best place, should this be
// seperated? When editing an issue the password is not required for the first
// contact but should be requested and used so avoid wasting users time.
func makeRequest(c Config, json []byte) (Reply, error) {

	var reply Reply
	// Set the correct url for the request.
	c, addr, err := setURL(c)
	if err != nil {
		return reply, fmt.Errorf("setURL: %v", err)
	}

	// Compose an array of header key value pairs.
	addr.header, err = composeHeader(c)
	if err != nil {
		return reply, fmt.Errorf("composeHeader: %v", err)
	}

	if f&cVERBOSE > 0 {
		fmt.Printf("makeRequest: json %v\n", string(json))
	}

	buf := bytes.NewReader(json)

	resp, err := sendRequest(c, addr, buf)
	if err != nil {
		return reply, fmt.Errorf("sendRequest: %v", err)
	}

	// Record the responce return status.
	addr.Status, err = getStatus(resp)
	if err != nil {
		return reply, fmt.Errorf("getStatus: %v", err)
	}

	if f&cVERBOSE > 0 {
		fmt.Printf("getStatus: %v\n", addr.String())
	}

	// Decode the responce if required.
	if f&(cLIST|cREAD) > 0 {
		reply, err = respDecode(c, resp)
		if err != nil {
			return reply, fmt.Errorf("respDecode: %v", err)
		}
	}
	resp.Body.Close()

	return reply, err
}
