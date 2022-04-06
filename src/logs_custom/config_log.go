package logs_custom

import (
	"github.com/kataras/golog"
	"os"
	"time"
)

var logger *golog.Logger

func todayFilename(preFixLink string) string {
	today := preFixLink + "Iris_" + time.Now().Format("2006_01_02")
	return today + ".log"
}

func NewLogFile(preFixLink string) *os.File {
	filename := todayFilename(preFixLink)
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func SetLogger(log *golog.Logger) {
	logger = log
}

func Logger() *golog.Logger {
	if logger == nil {
		panic("please init logs before use")
	}

	return logger
}
