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
func basicRequest(h []Header) []Header {

	// Set header.
	h = append(h, accept())

	return h
}

// authRequest generates a request header that uses oauth2 authorisation.
func authRequest(c Config, h []Header) []Header {

	// If token provided use that, else request password.
	if f&cTOKEN > 0 {
		h = append(h, authorize(c))
	} else {
		h = append(h, password(c))
	}

	return h
}

// composeHeader uses the current confiuration to set the correct header for
// the required HTTP request.
// TODO NOW define constants
func composeHeader(c Config) ([]Header, error) {

	var h []Header

	// Set basic request
	h = basicRequest(h)

	// Set autrorisation details.
	if f&cAUTH > 0 || f&cTOKEN > 0 {
		h = authRequest(c, h)
	}

	return h, nil
}
