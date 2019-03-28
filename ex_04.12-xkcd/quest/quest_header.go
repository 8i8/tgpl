package quest

// accept defines the media type.
func accAppJson() Header {
	var h Header
	h.Key = "Accept"
	h.Value = "application/json"
	return h
}

func accEncGzip() Header {
	var h Header
	h.Key = "Accept-Encoding"
	h.Value = "gzip, deflate, br"
	return h
}
