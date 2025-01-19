//go:build android
// +build android

package zip7

import (
	"os"
	"path/filepath"
)

// 使用绝对路径访问可执行文件 7za
var _BINARY string

func init() {
	exe, _ := os.Executable()
	_BINARY = filepath.Join(filepath.Dir(exe), "./7za")
}
