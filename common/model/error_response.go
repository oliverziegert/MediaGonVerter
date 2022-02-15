package model

// Error information
type ErrorResponse struct {
	// HTTP status code
	Code int32 `json:"code"`
	// HTTP status code description
	Message string `json:"message"`
	// Debug information
	DebugInfo string `json:"debugInfo,omitempty"`
	// Internal error code
	ErrorCode int32 `json:"errorCode,omitempty"`
}
