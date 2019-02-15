package github

// accept is a standtard header line.
func accept(c Config) Header {
	var h Header
	h.Key = "Accept"
	h.Value = "application/vnd.github.v3.text-match+json"
	return h
}

// authorize is the standard header Oath2 authorisation.
func authorize(conf Config) Header {
	var h Header
	h.Key = "Authorization"
	h.Value = "token " + conf.Token
	return h
}

// basicRequest generates the most basic hearder for the program.
func basicRequest(c Config, h []Header) []Header {

	// Add header to request.
	h = append(h, accept(c))

	return h
}

// authRequest generates a request header that uses oauth2 authorisation.
func authRequest(c Config, h []Header) []Header {

	// Set header.
	h = append(h, accept(c))
	h = append(h, authorize(c))

	return h
}

// composeHeader uses the current confiuration to set the correct header for
// the required HTTP request.
func composeHeader(c Config) ([]Header, error) {

	var h []Header

	switch rState {
	case rMany:
		h = basicRequest(c, h)
	case rLone:
		h = basicRequest(c, h)
	case rNone:
		h = authRequest(c, h)
	}
	return h, nil
}
