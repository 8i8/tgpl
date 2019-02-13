package gitish

func accept(c Config) Header {
	var h Header
	h.Key = "Accept"
	h.Value = "application/vnd.github.v3.text-match+json"
	return h
}

func authorize(conf Config) Header {
	var h Header
	h.Key = "Authorization"
	h.Value = "token " + conf.Token
	return h
}

// SearchIssues queries the GitHub issue tracker.
func basicRequest(c Config, h []Header) []Header {

	// Add header to request.
	h = append(h, accept(c))

	return h
}

// Generate a new issue.
func authRequest(c Config, h []Header) []Header {

	// Set header.
	h = append(h, accept(c))
	h = append(h, authorize(c))

	return h
}

func composeHeader(c Config) ([]Header, error) {

	var h []Header

	// Get header.
	switch state {
	case rMany:
		h = basicRequest(c, h)
	case rLone:
		h = basicRequest(c, h)
	case rNone:
		h = authRequest(c, h)
	}
	return h, nil
}
