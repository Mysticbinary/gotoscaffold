package gtserr

import "os"

// 统一的错误处理，有错误就退出程序
func ExitOnError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
