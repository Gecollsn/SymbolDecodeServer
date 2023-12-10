package tools

import (
	"path/filepath"
	"runtime"
)

func Dirname() string {
	_, caller, _, _ := runtime.Caller(1)
	return filepath.Dir(caller)
}
