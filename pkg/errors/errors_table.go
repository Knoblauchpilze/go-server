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

	ErrFailedToGetBody
	ErrBodyParsingFailed

	ErrNoSuchHeader
	ErrNonUniqueHeader

	ErrNoResponse
	ErrResponseIsError

	ErrInvalidSignUpData
	ErrPostRequestFailed
)

var errorsCodeToMessage = map[ErrorCode]string{
	ErrInvalidUserName:     "user name is invalid",
	ErrInvalidPassword:     "password is invalid",
	ErrUserAlreadyExists:   "user already exists",
	ErrUserCreationFailure: "internal error while creating user",
	ErrNoSuchUser:          "no such user",

	ErrNoSuchToken:        "no such token",
	ErrTokenAlreadyExists: "token already exists",

	ErrFailedToGetBody:   "failed to get request body",
	ErrBodyParsingFailed: "failed to parse request body",

	ErrNoSuchHeader:    "no such header in request",
	ErrNonUniqueHeader: "header is defined multiple times in request",

	ErrNoResponse:      "no response",
	ErrResponseIsError: "response returned error code",

	ErrInvalidSignUpData: "invalid sign up data",
	ErrPostRequestFailed: "post request failed",
}

var defaultErrorMessage = "unexpected error occurred"
