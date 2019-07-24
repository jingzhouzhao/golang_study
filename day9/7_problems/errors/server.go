package errors

import (
	"errors"
)

var (
	ListenError = errors.New("Server listen error")
	HeaderByteError = errors.New("Header byte error")
	UnknownCmd = errors.New("Unknown Cmd")
)

