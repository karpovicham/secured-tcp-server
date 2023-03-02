package proto

type ErrorCode int

const (
	// ErrorCodeUnavailable - Request data is not valid (not correct data format for the Type)
	ErrorCodeUnavailable ErrorCode = iota

	// ErrorCodeInvalidData - Request data is not valid (not correct data format for the Type)
	ErrorCodeInvalidData

	// ErrorCodeInvalidParameter - Some parameter is not valid the request
	ErrorCodeInvalidParameter

	// ErrorCodeUnknownRequestType - Request type is not supported
	ErrorCodeUnknownRequestType

	// ErrorCodeInvalidCredentials - Credentials are not valid (incorrect password)
	ErrorCodeInvalidCredentials

	// ErrorCodeUnauthenticated - Resource access requires the authentication (login)
	ErrorCodeUnauthenticated

	// ErrorCodeNotFound - Resource not found
	ErrorCodeNotFound
)

// easyjson:json
type ErrorResponseData struct {
	Code ErrorCode
	// Error context
	Msg string
}
