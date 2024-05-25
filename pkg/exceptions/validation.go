package exceptions

import "fmt"

type Validation struct {
	Code    int
	Message string
}

func (ex Validation) Error() string {
	return fmt.Sprintf("error %d: %s", ex.Code, ex.Message)
}
