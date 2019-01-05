package code

const (
	ResultOK                  = "000"
	ResultResetContent        = "Reset Content"
	ResultInternalServerError = "Internal Server Error"
	ResultNotFoundDataInRedis = "Not Found Data In Redis"

	ResultInvalidAccessToken          = "Invalid access token"
	ResultInvalidApplicationKeySecret = "Invalid application key/secret"

	ResultDBIsNotInitialized      = "DB is not initialized"
	ResultStorageIsNotInitialized = "Storage is not initialized"
	ResultApplicationContextIsNil = "Application context is nil"

	ResultInvalidAPIVersion                 = "Invalid API version"
	ResultUserIDIsRequired                  = "user_id is required"
	ResultUserTypeIsRequired                = "user_type is required"
	ResultServiceIDIsRequired               = "service_id is required"
	ResultServiceTypeIsRequired             = "service_type is required"
	ResultApplicationKeyIsRequired          = "Application-Key is required"
	ResultApplicationSecretIsRequired       = "Application-Secret is required"
	ResultServerApplicationKeyIsRequired    = "Server-Application-Key is required"
	ResultServerApplicationSecretIsRequired = "Server-Application-Secret is required"
	ResultClientApplicationKeyIsRequired    = "Client-Application-Key is required"
	ResultClientApplicationSecretIsRequired = "Client-Application-Secret is required"
	ResultApplicationNameIsRequired         = "application_name is required"
	ResultApplicationDescriptionIsRequired  = "application_description is required"
	ResultInvalidParameter                  = "invalid parameter"

	ResultPermissionDenied       = "Permission denied"
	ResultClientPermissionDenied = "Client permission denied"

	ResultApplicationNotFound       = "Application is not found"
	ResultClientApplicationNotFound = "Client application is not found"
	ResultServerApplicationNotFound = "Server application is not found"
	ResultApplicationAlreadyExists  = "Applicationalready exists"
	ResultPermissionAlreadyExists   = "Permission already exists"
)
