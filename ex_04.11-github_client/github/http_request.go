package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	Http, Url string
}

// Header contains header request key paires.
type Header struct {
	Key, Value string
}

// Status contains an http responce status, both the text and the code.
type Status struct {
	Code    int
	Message string
}

func sendRequest(conf Config, addr Address, json io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(addr.Http, addr.Url, json)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %v", err)
	}

	headers, err := composeHeader(conf)
	if err != nil {
		return nil, fmt.Errorf("composeHeader: %v", err)
	}

	for _, h := range headers {
		req.Header.Set(h.Key, h.Value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient: %v", err)
	}
	return resp, err
}

func getStatus(resp *http.Response) (*Status, error) {

	s := &Status{}
	s.Code = resp.StatusCode
	s.Message = http.StatusText(resp.StatusCode)

	if (state == rLone || state == rMany) && s.Code != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("response: %v %v", s.Code, s.Message)
	} else if state == rNone && s.Code != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("response: %v %v", s.Code, s.Message)
	}

	return s, nil
}

func respDecode(c Config, resp *http.Response) (interface{}, error) {

	var err error
	// Decode reply urlAddr for a direct http request and urlSear using the
	// API search function.
	if state == rLone {
		if c.Verbose {
			fmt.Println("respDecode: single issue decode")
		}
		var issue Issue
		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		return issue, err

	} else if state == rMany {
		if c.Verbose {
			fmt.Println("respDecode: multiple issue decode")
		}
		var issue IssuesSearchResult
		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		result := issue.Items
		return result, err
	}

	return nil, nil
}

func treatResponce(c Config, I interface{}) error {

	var err error
	switch v := I.(type) {
	case []*Issue:
		err = listIssues(c, v)
	case Issue:
		printIssue(v)
	default:
		err = fmt.Errorf("unknown type")
	}
	return err
}

// MakeRequest orchestrates an http request.
func MakeRequest(conf Config, json io.Reader) (interface{}, error) {

	addr, err := setURL(conf)
	if err != nil {
		return nil, fmt.Errorf("setUrl failed: %v", err)
	}

	resp, err := sendRequest(conf, addr, json)
	if err != nil {
		return nil, fmt.Errorf("sendRequest: %v", err)
	}

	_, err = getStatus(resp)
	if err != nil {
		return nil, fmt.Errorf("getStatus: %v", err)
	}

	result, err := respDecode(conf, resp)
	if err != nil {
		return nil, fmt.Errorf("respDecode: %v", err)
	}

	resp.Body.Close()

	err = treatResponce(conf, result)

	return result, err
}
