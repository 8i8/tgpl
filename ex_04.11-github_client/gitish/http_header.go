package gitish

import "fmt"

// accept defines the media type.
func accept() Header {
	var h Header
	h.Key = "Accept"
	h.Value = "application/vnd.github.v3.text-match+json"
	return h
}

// authorize for Oath2 authorisation.
func authorize(c Config) Header {
	var h Header
	h.Key = "Authorization"
	h.Value = "token " + c.Token
	return h
}

// password creates a password request when OAuth2 is not being used.
func password(c Config) (Header, error) {
	var h Header
	h.Key = "Authorization"
	pass, err := getPass(c)
	if err != nil {
		return h, fmt.Errorf("password: getPass failed")
	}
	h.Value = "Basic " + pass
	return h, nil
}

// lock sets the reason for locking an issue, if one has been provided.
func lock(c Config) Header {
	var h Header
	h.Key = "active_lock_reason"
	h.Value = c.Reason
	return h
}

// authRequest seeks authorisation when it is required.
func authRequest(c Config, h []Header) ([]Header, error) {

	// If a token has been provided use that, else request a password.
	if f&cTOKEN > 0 {
		h = append(h, authorize(c))
	} else {
		pass, err := password(c)
		if err != nil {
			return nil, fmt.Errorf("authRequest: %v", err)
		}
		h = append(h, pass)
	}

	return h, nil
}

// composeHeader uses the current configuration to set the correct header for
// the required HTTP request.
func composeHeader(c Config) ([]Header, error) {

	var h []Header
	var err error

	// Set basic request
	h = append(h, accept())

	// Set authorisation details.
	if f&cAUTH > 0 {
		h, err = authRequest(c, h)
		if err != nil {
			return nil, err
		}
	}

	// Set Lock details.
	if f&cLOCK > 0 {
		h = append(h, lock(c))
	}

	return h, nil
}
