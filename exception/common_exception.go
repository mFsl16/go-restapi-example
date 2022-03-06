package exception

type CommonException struct {
	Error string
}

func NewCommonException(error string) CommonException {
	return CommonException{
		Error: error,
	}
}
