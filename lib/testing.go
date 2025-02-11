package lib

import (
	"path/filepath"
	"runtime"
)


func loadConfig() (*Config, error) {
	return NewConfig(getLocalPath("../tests/config.json"))
}


func getLocalPath(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), file)
}
