package errors

import (
	"errors"
)

var (
	ParameterError = errors.New("Parameter error")
	UserNotExists = errors.New("User not exists")
	PasswordError = errors.New("Password error")
	SerializeError = errors.New("Serialize error")
	DeserializeError = errors.New("Deserialize error")
)