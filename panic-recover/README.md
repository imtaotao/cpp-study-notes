`Golang` 没有结构化异常，使用 `panic` 抛出错误，`recover` 捕获错误。

panic:
1. 内置函数
2. 假如函数 `F` 中书写了 `panic` 语句，会终止其后要执行的代码，在 `panic` 所在函数 `F` 内如果存在要执行的 `defer` 函数列表，按照 `defer` 的逆序执行
3. 返回函数 `F` 的调用者 `G`，在 `G` 中，调用函数 `F` 语句之后的代码不会执行，假如函数 `G` 中存在要执行的 `defer` 函数列表，按照 `defer` 的逆序执行
4. 直到 `goroutine` 整个退出，并报告错误

recover:
1. 内置函数
2. 用来控制一个 `goroutine` 的 `panicking` 行为，捕获 `panic`，从而影响应用的行为
3. 一般的调用建议
    - 在 `defer` 函数中，通过 `recover` 来终止一个 `goroutine` 的 `panicking` 过程，从而恢复正常代码的执行
    - 可以获取通过 panic 传递的 error

注意:
1. 利用 `recover` 处理 `panic` 指令，`defer` 必须放在 `panic` 之前定义，另外 `recover` 只有在 `defer` 调用的函数中才有效。
  否则当 `panic` 时，`recover` 无法捕获到 `panic`，无法防止 `panic` 扩散。
2. `recover` 处理异常后，逻辑并不会恢复到 `panic` 那个点去，函数跑到 `defer` 之后的那个点。
3. 多个 `defer` 会形成 `defer` 栈，后定义的 `defer` 语句会被最先调用。
4. 由于 `panic`、`recover` 参数类型为 `interface{}`，因此可抛出任何类型对象。