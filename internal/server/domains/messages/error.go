package messages

const (
	// Repository Errors
	ErrFetchFailed    = "messages.fetch_failed"
	ErrCountFailed    = "messages.count_failed"
	ErrLoginFailed    = "messages.login_failed"
	ErrCreateFailed   = "messages.create_failed"
	ErrUpdateFailed   = "messages.update_failed"
	ErrDeleteFailed   = "messages.delete_failed"
	ErrIsExistsFailed = "messages.exists_check_failed"
	ErrNotFound       = "messages.not_found"

	// Domain Errors
	ErrSenderIDRequired = "messages.sender_id_required"
	ErrRoomIDRequired   = "messages.room_id_required"

	// Service Errors
	ErrMessageEncryptFailed = "messages.message_encrypt_failed"
	ErrMessageDecryptFailed = "messages.message_dencrypt_failed"
)
