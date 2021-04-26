// Package dsblock
// Author: Zoueature
// Email: zoueature@gmail.com
// -------------------------------

package dsblock

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path/filepath"
	"sync"
	"time"
)

type zookeeperLock struct {
	lockPath  string
	lockNode  string
	zk        *zk.Conn
	localLock sync.Mutex
	lock      string
}

func (z *zookeeperLock) fullPath(leafNode string) string {
	return z.lockPath + "/" + leafNode
}
func (z *zookeeperLock) Lock() error {
	// 本地加上互斥锁
	z.localLock.Lock()
	println("get local lock")

	// 创建临时有序节点
	nodeName, err := z.zk.Create(z.fullPath(z.lockNode), nil, zk.FlagSequence|zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	defer func() {
		z.lock = nodeName
	}()
	children, _, err := z.zk.Children(z.lockPath)
	fmt.Printf("%+v", children)
	if err != nil {
		return err
	}
	leafNode := getLeafNode(nodeName)
	if children[0] == leafNode {
		return nil
	}
	n := binarySearcher(children).Search(leafNode)
	watchKey := z.fullPath(children[n-1])
	_, _, nodeChan, err := z.zk.GetW(watchKey)
	if err != nil {
		return err
	}
	exist, _, err := z.zk.Exists(watchKey)
	if err != nil {
		return err
	}
	if !exist {
		println("get distribute lock " + nodeName)
		return nil
	}
	for {
		event := <-nodeChan
		if event.Type == zk.EventNodeDeleted {
			// 节点删除， 获得锁
			println(watchKey + " deleted")
			break
		}
	}
	println("get distribute lock: " + nodeName)
	return nil
}

func (z *zookeeperLock) UnLock() error {
	err := z.zk.Delete(z.lock, -1)
	if err != nil {
		return err
	}
	z.localLock.Unlock()
	return nil
}

func (z *zookeeperLock) Close() {
	z.zk.Close()
}

var zkLocker DsbLock

func ZookeeperLocker(host []string, lockPath, lockNode string) DsbLock {
	if zkLocker == nil {
		conn, _, err := zk.Connect(host, 5*time.Second)
		if err != nil {
			panic(fmt.Sprintf("connect zk error: %s", err.Error()))
		}
		zkLocker = &zookeeperLock{
			lockPath:  lockPath,
			lockNode:  lockNode,
			zk:        conn,
			localLock: sync.Mutex{},
		}
	}
	return zkLocker
}

type binarySearcher []string

func (receiver binarySearcher) Search(key string) int {
	n := len(receiver)
	if n == 0 {
		return -1
	}
	i, j := 0, n-1
	for {
		k := (i + j) / 2
		v := receiver[k]
		if v == key {
			return k
		}
		if i == j {
			break
		}
		if v > key {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	return -1
}

func getLeafNode(fullPath string) string {
	_, node := filepath.Split(fullPath)
	return node
}
