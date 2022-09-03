package some_utils

import (
	"os"
	"strconv"
)

func GetUintEnv(key string) (uint, error) {
	ui, err := strconv.ParseUint(os.Getenv(key), 10, 64)

	return uint(ui), err
}
