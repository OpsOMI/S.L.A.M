package commons

const (
	// HTTP Status Codes
	StatusOK                  = "OK"
	StatusBadRequest          = "BadRequest"
	StatusUnauthorized        = "Unauthorized"
	StatusForbidden           = "Forbidden"
	StatusNotFound            = "NotFound"
	StatusConflict            = "Conflict"
	StatusTooManyRequests     = "TooManyRequests"
	StatusInternalServerError = "InternalServerError"
	StatusServiceUnavailable  = "ServiceUnavailable"

	// Response Types
	ResponseTypeSuccess = "Success"
	ResponseTypeError   = "Error"

	// General Errors
	ErrInvalidID = "general.invalid_id"

	// Roles
	RoleOwner = "owner"
	RoleUser  = "user"

	// File System Errors
	ErrFileCreateDir = "file.create_dir_error"
	ErrFileCreate    = "file.create_error"
	ErrFileDelete    = "file.delete_error"
	ErrFileCopy      = "file.copy_error"
	ErrFileOpen      = "file.open_error"

	// Parse
	ErrParseFailed = "parse.invalid_json_payload"
)
