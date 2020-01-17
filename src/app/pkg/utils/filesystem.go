package utils

import (
	"io/ioutil"
	"os"
)

func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadFile(filepath string) string {
	file, err := os.Open(filepath); Check(err)
	data, err := ioutil.ReadAll(file); Check(err)
	return string(data)
}
