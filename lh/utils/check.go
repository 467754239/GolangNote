package utils

import (
	"os"

	"github.com/467754239/GolangNote/lh/utils/log"
)

func checkError(e error, stage string) {
	if e != nil {
		log.Error("[err-%s] %v", stage, e)
		os.Exit(1)
	}
}
