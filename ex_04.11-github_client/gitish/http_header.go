package gitish

// accept is a standard header line.
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

// password requests and creates a basic password login.
func password(c Config) Header {
	var h Header
	h.Key = "Authorization"
	pass, _ := getPass(c)
	h.Value = "Basic " + pass
	return h
}

// lock: set a reason for locking an issue.
func lock(c Config) Header {
	var h Header
	h.Key = "active_lock_reason"
	h.Value = c.Reason
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

// composeHeader uses the current configuration to set the correct header for
// the required HTTP request.
func composeHeader(c Config) ([]Header, error) {

	var h []Header

	// Set basic request
	h = append(h, accept())

	// Set authorisation details.
	if f&cAUTH > 0 {
		h = authRequest(c, h)
	}

	// Set Lock details.
	// if f&cLOCK > 0 {
	// 	h = append(h, lock(c))
	// }

	return h, nil
}
