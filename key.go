package dsblock

import "fmt"

type LockKey string

func (o LockKey) String(args ...interface{}) string {
	return fmt.Sprintf(string(o), args...)
}
