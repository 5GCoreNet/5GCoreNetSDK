package fivegc

// StatusCode represents the status code of the HTTP response.
type StatusCode uint16

const (
	StatusContinue StatusCode = 100 // RFC 9110, 15.2.1

	StatusSwitchingProtocols StatusCode = 101 // RFC 9110, 15.2.2

	StatusProcessing StatusCode = 102 // RFC 2518, 10.1

	StatusEarlyHints StatusCode = 103 // RFC 8297

	StatusOK StatusCode = 200 // RFC 9110, 15.3.1

	StatusCreated StatusCode = 201 // RFC 9110, 15.3.2

	StatusAccepted StatusCode = 202 // RFC 9110, 15.3.3

	StatusNonAuthoritativeInfo StatusCode = 203 // RFC 9110, 15.3.4

	StatusNoContent StatusCode = 204 // RFC 9110, 15.3.5

	StatusResetContent StatusCode = 205 // RFC 9110, 15.3.6

	StatusPartialContent StatusCode = 206 // RFC 9110, 15.3.7

	StatusMultiStatus StatusCode = 207 // RFC 4918, 11.1

	StatusAlreadyReported StatusCode = 208 // RFC 5842, 7.1

	StatusIMUsed StatusCode = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices StatusCode = 300 // RFC 9110, 15.4.1

	StatusMovedPermanently StatusCode = 301 // RFC 9110, 15.4.2

	StatusFound StatusCode = 302 // RFC 9110, 15.4.3

	StatusSeeOther StatusCode = 303 // RFC 9110, 15.4.4

	StatusNotModified StatusCode = 304 // RFC 9110, 15.4.5

	StatusUseProxy StatusCode = 305 // RFC 9110, 15.4.6

	_ StatusCode = 306 // RFC 9110, 15.4.7 (Unused)

	StatusTemporaryRedirect StatusCode = 307 // RFC 9110, 15.4.8

	StatusPermanentRedirect StatusCode = 308 // RFC 9110, 15.4.9

	StatusBadRequest StatusCode = 400 // RFC 9110, 15.5.1

	StatusUnauthorized StatusCode = 401 // RFC 9110, 15.5.2

	StatusPaymentRequired StatusCode = 402 // RFC 9110, 15.5.3

	StatusForbidden StatusCode = 403 // RFC 9110, 15.5.4

	StatusNotFound StatusCode = 404 // RFC 9110, 15.5.5

	StatusMethodNotAllowed StatusCode = 405 // RFC 9110, 15.5.6

	StatusNotAcceptable StatusCode = 406 // RFC 9110, 15.5.7

	StatusProxyAuthRequired StatusCode = 407 // RFC 9110, 15.5.8

	StatusRequestTimeout StatusCode = 408 // RFC 9110, 15.5.9

	StatusConflict StatusCode = 409 // RFC 9110, 15.5.10

	StatusGone StatusCode = 410 // RFC 9110, 15.5.11

	StatusLengthRequired StatusCode = 411 // RFC 9110, 15.5.12

	StatusPreconditionFailed StatusCode = 412 // RFC 9110, 15.5.13

	StatusRequestEntityTooLarge StatusCode = 413 // RFC 9110, 15.5.14

	StatusRequestURITooLong StatusCode = 414 // RFC 9110, 15.5.15

	StatusUnsupportedMediaType StatusCode = 415 // RFC 9110, 15.5.16

	StatusRequestedRangeNotSatisfiable StatusCode = 416 // RFC 9110, 15.5.17

	StatusExpectationFailed StatusCode = 417 // RFC 9110, 15.5.18

	StatusTeapot StatusCode = 418 // RFC 9110, 15.5.19 (Unused)

	StatusMisdirectedRequest StatusCode = 421 // RFC 9110, 15.5.20

	StatusUnprocessableEntity StatusCode = 422 // RFC 9110, 15.5.21

	StatusLocked StatusCode = 423 // RFC 4918, 11.3

	StatusFailedDependency StatusCode = 424 // RFC 4918, 11.4

	StatusTooEarly StatusCode = 425 // RFC 8470, 5.2.

	StatusUpgradeRequired StatusCode = 426 // RFC 9110, 15.5.22

	StatusPreconditionRequired StatusCode = 428 // RFC 6585, 3

	StatusTooManyRequests StatusCode = 429 // RFC 6585, 4

	StatusRequestHeaderFieldsTooLarge StatusCode = 431 // RFC 6585, 5

	StatusUnavailableForLegalReasons StatusCode = 451 // RFC 7725, 3

	StatusInternalServerError StatusCode = 500 // RFC 9110, 15.6.1

	StatusNotImplemented StatusCode = 501 // RFC 9110, 15.6.2

	StatusBadGateway StatusCode = 502 // RFC 9110, 15.6.3

	StatusServiceUnavailable StatusCode = 503 // RFC 9110, 15.6.4

	StatusGatewayTimeout StatusCode = 504 // RFC 9110, 15.6.5

	StatusHTTPVersionNotSupported StatusCode = 505 // RFC 9110, 15.6.6

	StatusVariantAlsoNegotiates StatusCode = 506 // RFC 2295, 8.1

	StatusInsufficientStorage StatusCode = 507 // RFC 4918, 11.5

	StatusLoopDetected StatusCode = 508 // RFC 5842, 7.2

	StatusNotExtended StatusCode = 510 // RFC 2774, 7

	StatusNetworkAuthenticationRequired StatusCode = 511 // RFC 6585, 6

)

// StatusText returns a text for the HTTP status code. It returns the empty

// string if the code is unknown.

func StatusText(code StatusCode) string {

	switch code {

	case StatusContinue:

		return "Continue"

	case StatusSwitchingProtocols:

		return "Switching Protocols"

	case StatusProcessing:

		return "Processing"

	case StatusEarlyHints:

		return "Early Hints"

	case StatusOK:

		return "OK"

	case StatusCreated:

		return "Created"

	case StatusAccepted:

		return "Accepted"

	case StatusNonAuthoritativeInfo:

		return "Non-Authoritative Information"

	case StatusNoContent:

		return "No Content"

	case StatusResetContent:

		return "Reset Content"

	case StatusPartialContent:

		return "Partial Content"

	case StatusMultiStatus:

		return "Multi-Status"

	case StatusAlreadyReported:

		return "Already Reported"

	case StatusIMUsed:

		return "IM Used"

	case StatusMultipleChoices:

		return "Multiple Choices"

	case StatusMovedPermanently:

		return "Moved Permanently"

	case StatusFound:

		return "Found"

	case StatusSeeOther:

		return "See Other"

	case StatusNotModified:

		return "Not Modified"

	case StatusUseProxy:

		return "Use Proxy"

	case StatusTemporaryRedirect:

		return "Temporary Redirect"

	case StatusPermanentRedirect:

		return "Permanent Redirect"

	case StatusBadRequest:

		return "Bad Request"

	case StatusUnauthorized:

		return "Unauthorized"

	case StatusPaymentRequired:

		return "Payment Required"

	case StatusForbidden:

		return "Forbidden"

	case StatusNotFound:

		return "Not Found"

	case StatusMethodNotAllowed:

		return "Method Not Allowed"

	case StatusNotAcceptable:

		return "Not Acceptable"

	case StatusProxyAuthRequired:

		return "Proxy Authentication Required"

	case StatusRequestTimeout:

		return "Request Timeout"

	case StatusConflict:

		return "Conflict"

	case StatusGone:

		return "Gone"

	case StatusLengthRequired:

		return "Length Required"

	case StatusPreconditionFailed:

		return "Precondition Failed"

	case StatusRequestEntityTooLarge:

		return "Request Entity Too Large"

	case StatusRequestURITooLong:

		return "Request URI Too Long"

	case StatusUnsupportedMediaType:

		return "Unsupported Media Type"

	case StatusRequestedRangeNotSatisfiable:

		return "Requested Range Not Satisfiable"

	case StatusExpectationFailed:

		return "Expectation Failed"

	case StatusTeapot:

		return "I'm a teapot"

	case StatusMisdirectedRequest:

		return "Misdirected Request"

	case StatusUnprocessableEntity:

		return "Unprocessable Entity"

	case StatusLocked:

		return "Locked"

	case StatusFailedDependency:

		return "Failed Dependency"

	case StatusTooEarly:

		return "Too Early"

	case StatusUpgradeRequired:

		return "Upgrade Required"

	case StatusPreconditionRequired:

		return "Precondition Required"

	case StatusTooManyRequests:

		return "Too Many Requests"

	case StatusRequestHeaderFieldsTooLarge:

		return "Request Header Fields Too Large"

	case StatusUnavailableForLegalReasons:

		return "Unavailable For Legal Reasons"

	case StatusInternalServerError:

		return "Internal Server Error"

	case StatusNotImplemented:

		return "Not Implemented"

	case StatusBadGateway:

		return "Bad Gateway"

	case StatusServiceUnavailable:

		return "Service Unavailable"

	case StatusGatewayTimeout:

		return "Gateway Timeout"

	case StatusHTTPVersionNotSupported:

		return "HTTP Version Not Supported"

	case StatusVariantAlsoNegotiates:

		return "Variant Also Negotiates"

	case StatusInsufficientStorage:

		return "Insufficient Storage"

	case StatusLoopDetected:

		return "Loop Detected"

	case StatusNotExtended:

		return "Not Extended"

	case StatusNetworkAuthenticationRequired:

		return "Network Authentication Required"

	default:

		return ""

	}

}

// ToStatusCode converts the integer value to a status code
func ToStatusCode(status uint16) StatusCode {
	return StatusCode(status)
}

// ToInt converts the status code to an integer value
func (s StatusCode) ToInt() int {
	return int(s)
}
