package errors

// Error codes for different types of errors
const (
	// System errors (1000-1999)
	ErrInternalServer     = "ERR_INTERNAL_SERVER"
	ErrServiceUnavailable = "ERR_SERVICE_UNAVAILABLE"
	ErrDatabaseError      = "ERR_DATABASE_ERROR"
	ErrExternalService    = "ERR_EXTERNAL_SERVICE"

	// Validation errors (2000-2999)
	ErrInvalidInput     = "ERR_INVALID_INPUT"
	ErrMissingRequired  = "ERR_MISSING_REQUIRED"
	ErrInvalidFormat    = "ERR_INVALID_FORMAT"
	ErrValidationFailed = "ERR_VALIDATION_FAILED"

	// Authentication errors (3000-3999)
	ErrUnauthorized = "ERR_UNAUTHORIZED"
	ErrForbidden    = "ERR_FORBIDDEN"
	ErrInvalidToken = "ERR_INVALID_TOKEN"
	ErrTokenExpired = "ERR_TOKEN_EXPIRED"

	// Resource errors (4000-4999)
	ErrNotFound         = "ERR_NOT_FOUND"
	ErrAlreadyExists    = "ERR_ALREADY_EXISTS"
	ErrResourceConflict = "ERR_RESOURCE_CONFLICT"
	ErrResourceLocked   = "ERR_RESOURCE_LOCKED"

	// Business logic errors (5000-5999)
	ErrBusinessRule     = "ERR_BUSINESS_RULE"
	ErrInvalidOperation = "ERR_INVALID_OPERATION"
	ErrInsufficientData = "ERR_INSUFFICIENT_DATA"
	ErrInvalidState     = "ERR_INVALID_STATE"
)

// Predefined domain errors
var (
	// System errors
	InternalServerError     = NewDomainError(ErrInternalServer, "Internal server error")
	ServiceUnavailableError = NewDomainError(ErrServiceUnavailable, "Service temporarily unavailable")
	DatabaseError           = NewDomainError(ErrDatabaseError, "Database operation failed")
	ExternalServiceError    = NewDomainError(ErrExternalService, "External service error")

	// Validation errors
	InvalidInputError     = NewDomainError(ErrInvalidInput, "Invalid input provided")
	MissingRequiredError  = NewDomainError(ErrMissingRequired, "Required field is missing")
	InvalidFormatError    = NewDomainError(ErrInvalidFormat, "Invalid format provided")
	ValidationFailedError = NewDomainError(ErrValidationFailed, "Validation failed")

	// Authentication errors
	UnauthorizedError = NewDomainError(ErrUnauthorized, "Unauthorized access")
	ForbiddenError    = NewDomainError(ErrForbidden, "Access forbidden")
	InvalidTokenError = NewDomainError(ErrInvalidToken, "Invalid token")
	TokenExpiredError = NewDomainError(ErrTokenExpired, "Token has expired")

	// Resource errors
	NotFoundError         = NewDomainError(ErrNotFound, "Resource not found")
	AlreadyExistsError    = NewDomainError(ErrAlreadyExists, "Resource already exists")
	ResourceConflictError = NewDomainError(ErrResourceConflict, "Resource conflict")
	ResourceLockedError   = NewDomainError(ErrResourceLocked, "Resource is locked")

	// Business logic errors
	BusinessRuleError     = NewDomainError(ErrBusinessRule, "Business rule violation")
	InvalidOperationError = NewDomainError(ErrInvalidOperation, "Invalid operation")
	InsufficientDataError = NewDomainError(ErrInsufficientData, "Insufficient data")
	InvalidStateError     = NewDomainError(ErrInvalidState, "Invalid state")
)
