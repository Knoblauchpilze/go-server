package errors

type ErrorCode int

const (
	errGenericErrorCode ErrorCode = iota
	ErrInvalidUserName
	ErrInvalidPassword
	ErrUserAlreadyExists
	ErrUserCreationFailure
	ErrNoSuchUser

	ErrNoSuchToken
	ErrTokenAlreadyExists
)

var errorsCodeToMessage = map[ErrorCode]string{
	ErrInvalidUserName:     "user name is invalid",
	ErrInvalidPassword:     "password is invalid",
	ErrUserAlreadyExists:   "user already exists",
	ErrUserCreationFailure: "internal error while creating user",
	ErrNoSuchUser:          "no such user",

	ErrNoSuchToken:        "no such token",
	ErrTokenAlreadyExists: "token already exists",
}

var defaultErrorMessage = "unexpected error occurred"
