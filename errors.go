package main

type EmptyTokenError struct {
	message string
}

func (e EmptyTokenError) Error() string {
	if e.message != "" {
		return e.message
	} else {
		return "Received: Empty Token"
	}
}

type JWTParsingError struct {
	message string
}

func (e JWTParsingError) Error() string {
	if e.message != "" {
		return e.message
	} else {
		return "Error: Could not parse JWT token."
	}
}

type UnknownClaimTypeError struct {
	message string
}

func (e UnknownClaimTypeError) Error() string {
	if e.message != "" {
		return e.message
	} else {
		return "Error: Unknown claim type"
	}
}
