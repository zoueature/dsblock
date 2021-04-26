// Package dsblock
// Author: Zoueature
// Email: zoueature@gmail.com
// -------------------------------

package dsblock

type DsbLock interface {
	Lock() error
	UnLock() error
	Close()
}