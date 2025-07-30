package users

const (
	// Repository Errors
	ErrNotFound         = "userauth.not_found"
	ErrFetchFailed      = "userauth.fetch_failed"
	ErrCreateFailed     = "userauth.create_failed"
	ErrUpdateFailed     = "userauth.update_failed"
	ErrDeleteFailed     = "userauth.delete_failed"
	ErrSoftDeleteFailed = "userauth.soft_delete_failed"
	ErrCountFailed      = "userauth.count_failed"
	ErrIsExistsFailed   = "userauth.exists_check_failed"
	ErrFirstLoginFailed = "userauth.first_login_failed"
	ErrVerifyMailFailed = "userauth.verify_email_failed"

	// Domain Errors
	ErrNicknameRequired = "users.nickname_required"
	ErrNicknameTooLong  = "users.nickname_too_long"
	ErrNicknameTooShort = "users.nickname_too_short"

	ErrUsernameRequired = "users.username_required"
	ErrUsernameTooLong  = "users.username_too_long"
	ErrUsernameTooShort = "users.username_too_short"

	ErrPasswordRequired = "users.password_required"
)
