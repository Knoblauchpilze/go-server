package auth

import "fmt"

// auth
var ErrInvalidPassword = fmt.Errorf("password is invalid")
var ErrTokenAlreadyExists = fmt.Errorf("token already exists")
var ErrTokenCreationFailure = fmt.Errorf("internal error while creating token")
var ErrNoSuchToken = fmt.Errorf("no such token")
