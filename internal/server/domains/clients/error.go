package clients

const (
	// Repository Errors
	ErrFetchFailed    = "clients.fetch_failed"
	ErrCountFailed    = "clients.count_failed"
	ErrLoginFailed    = "clients.login_failed"
	ErrCreateFailed   = "clients.create_failed"
	ErrUpdateFailed   = "clients.update_failed"
	ErrDeleteFailed   = "clients.delete_failed"
	ErrIsExistsFailed = "clients.exists_check_failed"
	ErrNotFound       = "clients.not_found"

	// Domain Errors
	ErrUserIDRequired    = "clients.user_id_required"
	ErrClientKeyRequired = "clients.client_key_required"

	// Service Errors
	ErrConfigCreateFailed = "clients.config_create_failed"
	ErrTslCertCopyFailed  = "clients.tsl_server_copy_failed"
	ErrBuildClientFailed  = "clients.build_client_failed"
)
