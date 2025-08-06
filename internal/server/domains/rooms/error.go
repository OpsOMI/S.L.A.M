package rooms

const (
	// Repository Errors
	ErrFetchFailed    = "rooms.fetch_failed"
	ErrCountFailed    = "rooms.count_failed"
	ErrLoginFailed    = "rooms.login_failed"
	ErrCreateFailed   = "rooms.create_failed"
	ErrUpdateFailed   = "rooms.update_failed"
	ErrDeleteFailed   = "rooms.delete_failed"
	ErrIsExistsFailed = "rooms.exists_check_failed"
	ErrNotFound       = "rooms.not_found"

	// Domain Errors
	ErrOwnerIDRequired  = "rooms.owner_id_required"
	ErrRoomCodeRequired = "rooms.room_code_required"

	// Service Errors
	ErrPasswordHashFailed = "rooms.password_hash_failed"
	ErrInvalidCredentials = "rooms.invalid_credentials"
	ErrCodeGenerateFailed = "rooms.code_generate_failed"
)
