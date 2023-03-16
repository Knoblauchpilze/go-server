package errors

type ErrorCode int

const (
	errGenericErrorCode ErrorCode = iota
	ErrInvalidUserName
	ErrInvalidPassword
	ErrUserAlreadyExists
	ErrUserCreationFailure
	ErrNoSuchUser
)

var errorsCodeToMessage = map[ErrorCode]string{
	ErrInvalidUserName:     "user name is invalid",
	ErrInvalidPassword:     "password is invalid",
	ErrUserAlreadyExists:   "user already exists",
	ErrUserCreationFailure: "internal error while creating user",
	ErrNoSuchUser:          "no such user",
}

var defaultErrorMessage = "unexpected error occurred"
