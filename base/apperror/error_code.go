package apperror

type ErrCode int

const (
	ERR_INVALID_RESOURCE ErrCode = 1000
	ERR_INPUT_FORMAT     ErrCode = 1001
	ERR_INPUT_DATA       ErrCode = 1002
	ERR_NOT_FOUND                = 2000
	ERR_DATABASE         ErrCode = 4000
	ERR_UNKNOWN          ErrCode = 9000
)
