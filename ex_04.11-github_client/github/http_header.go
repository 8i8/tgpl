package github

// accept is a standtard header line.
func accept() Header {
	var h Header
	h.Key = "Accept"
	h.Value = "application/vnd.github.v3.text-match+json"
	return h
}

// authorize is the standard header Oath2 authorisation.
func authorize(c Config) Header {
	var h Header
	h.Key = "Authorization"
	h.Value = "token " + c.Token
	return h
}

// password requests and creats a basic password login.
func password(c Config) Header {
	var h Header
	h.Key = "Authorization"
	pass, _ := getPass(c)
	h.Value = "Basic " + pass
	return h
}

// basicRequest generates the most basic hearder for the program.
func basicRequest(c Config, h []Header) []Header {

	// Set header.
	h = append(h, accept())

	return h
}

// authRequest generates a request header that uses oauth2 authorisation.
func authRequest(c Config, h []Header) []Header {

	// Set header.
	h = append(h, accept())

	// If token provided use that, else request password.
	if len(c.Token) > 0 {
		h = append(h, authorize(c))
	} else {
		h = append(h, password(c))
	}

	return h
}

// Add token and authorise if token present.
func checkTokenRequest(c Config, h []Header) []Header {

	// Set header.
	h = append(h, accept())
	if len(c.Token) > 0 {
		h = append(h, authorize(c))
	}
	return h
}

// composeHeader uses the current confiuration to set the correct header for
// the required HTTP request.
// TODO NOW define constants
func composeHeader(c Config) ([]Header, error) {

	var h []Header

	switch {
	case f&cMANY > 0:
		h = basicRequest(c, h)
	case f&cLONE > 0:
		h = basicRequest(c, h)
	case f&cNONE > 0:
		h = authRequest(c, h)
	case f&cRAW > 0:
		h = checkTokenRequest(c, h)
	default:
		h = basicRequest(c, h)
	}
	return h, nil
}
