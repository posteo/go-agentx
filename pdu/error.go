// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import "fmt"

// The various pdu packet errors.
const (
	ErrorNone                  Error = 0
	ErrorOpenFailed            Error = 256
	ErrorNotOpen               Error = 257
	ErrorIndexWrongType        Error = 258
	ErrorIndexAlreadyAllocated Error = 259
	ErrorIndexNoneAvailable    Error = 260
	ErrorIndexNotAllocated     Error = 261
	ErrorUnsupportedContext    Error = 262
	ErrorDuplicateRegistration Error = 263
	ErrorUnknownRegistration   Error = 264
	ErrorUnknownAgentCaps      Error = 265
	ErrorParse                 Error = 266
	ErrorRequestDenied         Error = 267
	ErrorProcessing            Error = 268
)

// Error defines a pdu packet error.
type Error uint16

func (e Error) String() string {
	switch e {
	case ErrorNone:
		return "ErrorNone"
	case ErrorOpenFailed:
		return "ErrorOpenFailed"
	case ErrorNotOpen:
		return "ErrorNotOpen"
	case ErrorIndexWrongType:
		return "ErrorIndexWrongType"
	case ErrorIndexAlreadyAllocated:
		return "ErrorIndexAlreadyAllocated"
	case ErrorIndexNoneAvailable:
		return "ErrorIndexNoneAvailable"
	case ErrorIndexNotAllocated:
		return "ErrorIndexNotAllocated"
	case ErrorUnsupportedContext:
		return "ErrorUnsupportedContext"
	case ErrorDuplicateRegistration:
		return "ErrorDuplicateRegistration"
	case ErrorUnknownRegistration:
		return "ErrorUnknownRegistration"
	case ErrorUnknownAgentCaps:
		return "ErrorUnknownAgentCaps"
	case ErrorParse:
		return "ErrorParse"
	case ErrorRequestDenied:
		return "ErrorRequestDenied"
	case ErrorProcessing:
		return "ErrorProcessing"
	}
	return fmt.Sprintf("ErrorUnknown (%d)", e)
}
