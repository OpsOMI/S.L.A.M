package users

const (
	// Repository Errors
	ErrFetchFailed    = "users.fetch_failed"
	ErrCountFailed    = "users.count_failed"
	ErrLoginFailed    = "users.login_failed"
	ErrCreateFailed   = "users.create_failed"
	ErrUpdateFailed   = "users.update_failed"
	ErrDeleteFailed   = "users.delete_failed"
	ErrIsExistsFailed = "users.exists_check_failed"
	ErrNotFound       = "users.not_found"

	// Domain Errors
	ErrNicknameRequired = "users.nickname_required"
	ErrNicknameTooLong  = "users.nickname_too_long"
	ErrNicknameTooShort = "users.nickname_too_short"
	ErrUsernameRequired = "users.username_required"
	ErrUsernameTooLong  = "users.username_too_long"
	ErrUsernameTooShort = "users.username_too_short"
	ErrPasswordRequired = "users.password_required"
)
