package main

// 这是一个进阶知识点。
// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
// Go 语言中的 sync 包中提供了一个针对只执行一次场景的解决方案 – sync.Once。
// sync.Once 只有一个 Do 方法，其签名如下：
// `func (o *Once) Do(f func()) {}`

// 案例链接：http://www.topgoer.com/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/sync.html#synconce
