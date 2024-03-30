package logger

import (
	"fmt"
	"log"
	"os"
)

func Build(prefix string) *log.Logger {
	return log.New(
		os.Stdout,
		fmt.Sprintf("%s: ", prefix),
		log.Lmsgprefix|log.LstdFlags,
	)
}
