package service

import vpr "example"

func SetError(status int, message string) *vpr.Error {
	return &vpr.Error{Status: status, Message: message}
}
