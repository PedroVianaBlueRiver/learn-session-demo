package constants

type statusCode int

const (
	StatusContinue           statusCode = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols statusCode = 101 // RFC 9110, 15.2.2
	StatusProcessing         statusCode = 102 // RFC 2518, 10.1
	StatusEarlyHints         statusCode = 103 // RFC 8297

	StatusOK                   statusCode = 200 // RFC 9110, 15.3.1
	StatusCreated              statusCode = 201 // RFC 9110, 15.3.2
	StatusAccepted             statusCode = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo            = 203 // RFC 9110, 15.3.4
	StatusNoContent            statusCode = 204 // RFC 9110, 15.3.5
	StatusResetContent         statusCode = 205 // RFC 9110, 15.3.6
	StatusPartialContent       statusCode = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          statusCode = 207 // RFC 4918, 11.1
	StatusAlreadyReported      statusCode = 208 // RFC 5842, 7.1
	StatusIMUsed               statusCode = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  statusCode = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently statusCode = 301 // RFC 9110, 15.4.2
	StatusFound            statusCode = 302 // RFC 9110, 15.4.3
	StatusSeeOther         statusCode = 303 // RFC 9110, 15.4.4
	StatusNotModified      statusCode = 304 // RFC 9110, 15.4.5
	StatusUseProxy         statusCode = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect statusCode = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect statusCode = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   statusCode = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 statusCode = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              statusCode = 402 // RFC 9110, 15.5.3
	StatusForbidden                    statusCode = 403 // RFC 9110, 15.5.4
	StatusNotFound                     statusCode = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             statusCode = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                statusCode = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            statusCode = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               statusCode = 408 // RFC 9110, 15.5.9
	StatusConflict                     statusCode = 409 // RFC 9110, 15.5.10
	StatusGone                         statusCode = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               statusCode = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           statusCode = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        statusCode = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            statusCode = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         statusCode = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable statusCode = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            statusCode = 417 // RFC 9110, 15.5.18
	StatusTeapot                       statusCode = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           statusCode = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          statusCode = 422 // RFC 9110, 15.5.21
	StatusLocked                       statusCode = 423 // RFC 4918, 11.3
	StatusFailedDependency             statusCode = 424 // RFC 4918, 11.4
	StatusTooEarly                     statusCode = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              statusCode = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         statusCode = 428 // RFC 6585, 3
	StatusTooManyRequests              statusCode = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  statusCode = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   statusCode = 451 // RFC 7725, 3

	StatusInternalServerError           statusCode = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                statusCode = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    statusCode = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            statusCode = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                statusCode = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       statusCode = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         statusCode = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           statusCode = 507 // RFC 4918, 11.5
	StatusLoopDetected                  statusCode = 508 // RFC 5842, 7.2
	StatusNotExtended                   statusCode = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired statusCode = 511 // RFC 6585, 6
)
