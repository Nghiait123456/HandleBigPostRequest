package logsCustom

import (
	"os"
	"time"
)

func todayFilename() string {
	today := "log-file/" + time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func NewLogFile() *os.File {
	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
