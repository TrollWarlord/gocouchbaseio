package gocouchbaseio

import "fmt"

type agentError struct {
	message string
}

func (e agentError) Error() string {
	return e.message
}
func (e agentError) KeyNotFound() bool {
	return false
}
func (e agentError) KeyExists() bool {
	return false
}
func (e agentError) Timeout() bool {
	return false
}
func (e agentError) Temporary() bool {
	return false
}

type memdError struct {
	code statusCode
}

func (e memdError) Error() string {
	switch e.code {
	case success:
		return "Success."
	case keyNotFound:
		return "Key not found."
	case keyExists:
		return "Key already exists."
	case tooBig:
		return "Document value was too large."
	case notStored:
		return "The document could not be stored."
	case badDelta:
		return "An invalid delta was passed."
	case notMyVBucket:
		return "Operation sent to incorrect server."
	case unknownCommand:
		return "An unknown command was received."
	case outOfMemory:
		return "The server is out of memory."
	case tmpFail:
		return "A temporary failure occurred.  Try again later."
	default:
		return fmt.Sprintf("An unknown error occurred (%d).", e.code)
	}
}
func (e memdError) KeyNotFound() bool {
	return e.code == keyNotFound
}
func (e memdError) KeyExists() bool {
	return e.code == keyExists
}
func (e memdError) Timeout() bool {
	return false
}
func (e memdError) Temporary() bool {
	return e.code == outOfMemory || e.code == tmpFail
}
