package myUtils

import (
	"os"
	"strconv"
)

func getUintEnv(key string) (uint, error) {
	ui, err := strconv.ParseUint(os.Getenv(key), 10, 64)

	return uint(ui), err
}
