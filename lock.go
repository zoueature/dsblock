// Package dsblock
// Author: Zoueature
// Email: zoueature@gmail.com
// -------------------------------

package dsblock

import "time"

type Locker interface {
	Lock(key string, autoUnlockTime time.Duration) error
	UnLock(key string) error
	Close()
}
