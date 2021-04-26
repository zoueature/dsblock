// Package dsblock
// Author: Zoueature
// Email: zoueature@gmail.com
// -------------------------------

package dsblock

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path/filepath"
	"time"
)

type zookeeperLock struct {
	lockPath string
	lockNode string
	zk       *zk.Conn
}

func (z *zookeeperLock) fullPath(leafNode string) string {
	return z.lockPath + "/" + leafNode
}
func (z *zookeeperLock) Lock() error {
	// 创建临时有序节点
	nodeName, err := z.zk.Create(z.fullPath(z.lockNode), nil, zk.FlagSequence|zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	children, _, err := z.zk.Children(z.lockPath)
	if err != nil {
		return err
	}
	if children[0] == nodeName {
		return nil
	}
	n := binarySearcher(children).Search(leafNode(nodeName))
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
	return nil
}

func (z *zookeeperLock) UnLock() {
	panic("implement me")
}

func (z *zookeeperLock) Close() {
	z.zk.Close()
}

type locker func() DsbLock

func ZookeeperLocker(host []string, lockPath, lockNode string) DsbLock {
	conn, _, err := zk.Connect(host, 5 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("connect zk error: %s", err.Error()))
	}
	return &zookeeperLock{
		lockPath: lockPath,
		lockNode: lockNode,
		zk:       conn,
	}
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

func leafNode(fullPath string) string {
	_, node := filepath.Split(fullPath)
	return node
}