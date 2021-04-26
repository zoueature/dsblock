// Package dsblock
// Author: Zoueature           
// Email: zoueature@gmail.com  
// -------------------------------

package dsblock

import "testing"

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
    zl := ZookeeperLocker([]string{"127.0.0.1:2181"}, "/lock", "lock-")
    err := zl.Lock()
    if err != nil {
        t.Fatalf(err.Error())
    }
    zl.Close()
}

