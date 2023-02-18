package model

type ErrorResult struct {
	ErrorCode    int
	ErrorMessage string
}

func ErrorOf(errorCode int, errorMessage string) ErrorResult {
	return ErrorResult{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}
