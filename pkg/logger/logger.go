package logger

import (
	"io"
	"log"
	"os"
)

var Log *log.Logger

func Init(logFile string) error {
	var out io.Writer = os.Stdout
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		out = io.MultiWriter(os.Stdout, f)
	}
	Log = log.New(out, "", log.LstdFlags)
	return nil
}
