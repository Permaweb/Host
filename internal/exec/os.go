package exec

import (
	"os"
)

func rm(path string) (err error) {
	return os.RemoveAll(path)
}

func mv(first, second string) (err error) {
	return os.Rename(first, second)
}
