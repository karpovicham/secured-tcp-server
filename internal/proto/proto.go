// Package proto - describes messages structure for Client-Server TCP connection
package proto

// Message - Transferred data structure,
// Data Could be Request (from Client) or Response (from Server) structure related to the Type
// or the ErrorResponseData as the response in case of any error detected on the server
// easyjson:json
type Message struct {
	Type Type   `json:"type"`
	Data []byte `json:"data"`
}

// Type represents request/response type
// Client/Server apps should detect how to handle requests by this value
type Type int

const (
	// TypeError - General response message with Error code and context (could be returned to any request)
	TypeError Type = iota

	// TypeLogin - Request to authenticate user with Basic Auth, get user ID and session ID
	TypeLogin

	// TypeLogout - Request to log out the user
	// Requires Authentication
	TypeLogout

	// TypeModifyAccountSettings - Request to modify the user settings
	// Requires Authentication
	TypeModifyAccountSettings

	// TypeDeactivateUser - Request to deactivate the user (won't be able to log out anymore)
	// Requires Authentication
	TypeDeactivateUser

	// TypeAddFavoritePage - Request to Add user favorite page
	// Requires Authentication
	TypeAddFavoritePage
)

// Login request/response

// easyjson:json
type LoginRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// easyjson:json
type LoginResponseData struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	ExpireAt  int64  `json:"expire_at"`
}

// Logout request/response

// easyjson:json
type LogoutRequestData struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}

type LogoutResponseData struct{}

// Modify account settings request/response

// easyjson:json
type ModifyAccountSettingsRequestData struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	// Either NewUsername or NewEmail must be set
	NewUsername string `json:"new_username,omitempty"`
	NewEmail    string `json:"new_email,omitempty"`
}

type ModifyAccountSettingsResponseData struct{}

// Deactivate user request/response

// easyjson:json
type DeactivateUserRequestData struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}

type DeactivateUserResponseData struct{}

// Add favorite page to user request/response

// easyjson:json
type AddFavoritePageRequestData struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	PageURL   string `json:"page_url"`
}

type AddFavoritePageResponseData struct{}
