package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 前一个例子下的场景下就需要为 map 加锁来保证并发的安全性了，
// Go 语言的 sync 包中提供了一个开箱即用的并发安全版 map–sync.Map。
// 开箱即用表示不用像内置的 map 一样使用make函数初始化就能直接使用。
// 同时 sync.Map 内置了诸如 Store、Load、LoadOrStore、Delete、Range 等操作方法。

var m = sync.Map{} // 直接用 go 提供的这个就好

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
