package property

import "fmt"

type unsupportedRequestTypeError struct {
	Request interface{}
}

func (e unsupportedRequestTypeError) Error() string {
	return fmt.Sprintf("unsupported request type[%T]", e.Request)
}
