package kabusapi

import fmt "fmt"

type Error struct {
	Code           ErrorResponseCode
	Message        string
	HTTPStatusCode int
	HTTPStatus     string
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"%d %s: %d: %s",
		e.HTTPStatusCode,
		e.HTTPStatus,
		e.Code,
		e.Message,
	)
}
