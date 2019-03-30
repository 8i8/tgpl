package quest

import "strconv"

var VERBOSE bool

// Header key value pairs.
type Header struct {
	Key, Value string
}

// HttpQuest contains the details of an http request.
type HttpQuest struct {
	verb, url string
	header    []Header
	Msg       interface{}
	httpStatus
}

// GET sets the appropriate header HTTP verb and key pairs to connect by the
// GET method.
func (r *HttpQuest) GET(s string) {
	r.verb = "GET"
	r.url = s
	r.header = append(r.header, accAppJson())
}

// Status contains the error response code and description.
type httpStatus struct {
	Code    int
	Message string
}

// String method returns a string representation of the contained data.
func (s httpStatus) Status() string {
	return strconv.Itoa(s.Code) + " " + s.Message
}

// QuestGET GET from given url.
func (r HttpQuest) QuestGET(url string) (HttpQuest, error) {
	r.GET(url)
	return request(r, nil)
}
