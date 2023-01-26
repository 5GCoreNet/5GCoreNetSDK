package fivegc

// RedirectHeader represents the header of the HTTP response.
type RedirectHeader struct {
	Location  string `json:"Location"`
	SbiTarget string `json:"3gpp-Sbi-Target-Nf-Id"`
}

// RedirectResponse - The response shall include a Location header field containing a different URI  (pointing to a different URI of an other service instance), or the same URI if a request  is redirected to the same target resource via a different SCP.
type RedirectResponse struct {
	RedirectHeader RedirectHeader `json:"redirectHeader"`

	Cause string `json:"cause,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	TargetScp string `json:"targetScp,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	TargetSepp string `json:"targetSepp,omitempty"`
}
