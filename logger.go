package simple_logger

import (
	"fmt"
	"io"
	"time"
)

type LogLevel string

type Logger struct {
	OutputDest io.Writer
	ToConsole  bool
}

func InitLogger(outputDest io.Writer, toConsole bool) *Logger {
	return &Logger{
		OutputDest: outputDest,
		ToConsole:  toConsole,
	}
}

func (l *Logger) log(level LogLevel, message string) {
	currentTime := time.Now().Format("2006/01/24 15:04:05")
	message = fmt.Sprintf("[%s] %s, %s", string(level), currentTime, message)
	if l.ToConsole {
		fmt.Println(message)
	}
	if l.OutputDest != nil {
		fmt.Fprintln(l.OutputDest, message)
	}
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Warning(message string) {
	l.log("WARNING", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}
