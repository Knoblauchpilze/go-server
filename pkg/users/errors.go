package users

import "fmt"

// users
var ErrUserAlreadyExists = fmt.Errorf("user already exists")
var ErrInvalidUserName = fmt.Errorf("user name is invalid")
var ErrInvalidPassword = fmt.Errorf("password is invalid")
var ErrUserCreationFailure = fmt.Errorf("internal error while creating user")
var ErrNoSuchUser = fmt.Errorf("no such user")
