package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetCurrentDirectory() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("failed get current directory called from file is %q", filename)
	}
	return filepath.Dir(filename), nil
}
