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

	// SQL Transaction Errors
	ErrSQLTxCommit = "sql.tx_commit_error"
	ErrSQLTxCreate = "sql.tx_create_error"

	// File System Errors
	ErrFileInvalidType = "file.invalid_type"
	ErrFileCreateDir   = "file.create_dir_error"
	ErrFileCreate      = "file.create_error"
	ErrFileDelete      = "file.delete_error"
	ErrFileCopy        = "file.copy_error"
	ErrFileOpen        = "file.open_error"

	// Parse
	ErrParseFailed = "parse.invalid_json_payload"

	// Log Messages
	LogResendVerificationFailed = "Failed to resend verification email for [%s]: %v"
	LogSend2FACodeFailed        = "Failed to send 2FA code to [%s]: %v"
	LogSendPasswordResetFailed  = "Failed to send password reset email to [%s]: %v"
)
