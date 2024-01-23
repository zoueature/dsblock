// Package dsblock
// Author: Zoueature
// Email: zoueature@gmail.com
// -------------------------------

package zookeeper

import (
	"sync"
	"testing"
	"time"
)

func TestBinarySearcher_Search(t *testing.T) {
	value := binarySearcher{"1", "2", "3", "4"}
	n := value.Search("3")
	if n != 2 {
		t.Fatalf("not match")
	}
	value = binarySearcher{"1"}
	n = value.Search("3")
	if n != -1 {
		t.Fatalf("not match")
	}
	value = binarySearcher{"1", "1", "1", "1", "1", "1", "1", "1", "1"}
	n = value.Search("3")
	if n != -1 {
		t.Fatalf("not match")
	}
}

func TestZookeeperLock_Lock(t *testing.T) {
	zl := NewLocker([]string{"127.0.0.1:2181"}, "/lock", "lock-")
	wg := sync.WaitGroup{}
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := zl.Lock()
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			time.Sleep(5 * time.Second)
			err = zl.UnLock()
			if err != nil {
				t.Errorf(err.Error())
				return
			}
		}()
	}
	wg.Wait()
	zl.Close()
}
