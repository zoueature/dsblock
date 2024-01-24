package redis

import (
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	locker := NewSingleLocker("18.141.233.195:6379", "123456", 0)
	err := locker.Lock("123", 10*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	locker.UnLock("123")
	err = locker.Lock("123", 10*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
}
