/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

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
