package logs_custom

import (
	"os"
	"time"
)

func todayFilename(preFixLink string) string {
	today := preFixLink + time.Now().Format("Jan 02 2006")
	return today + ".txt"
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
