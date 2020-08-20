package kabusapi

import fmt "fmt"

type ErrorCode int32

type ErrorResponse struct {
	Code           ErrorCode
	Message        string
	HTTPStatusCode int
	HTTPStatus     string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%d %s: %d: %s",
		e.HTTPStatusCode,
		e.HTTPStatus,
		e.Code,
		e.Message,
	)
}
