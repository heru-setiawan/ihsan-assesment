package exceptions

import "fmt"

type Database struct {
	Code    int
	Message string
}

func (ex Database) Error() string {
	return fmt.Sprintf("error %d: %s", ex.Code, ex.Message)
}
