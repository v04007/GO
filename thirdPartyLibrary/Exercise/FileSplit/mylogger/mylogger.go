package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type logLevel uint64

const (
	UNKNOWN logLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (logLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("传入无效等级")
		return UNKNOWN, err
	}
}

func getLogString(s logLevel) string {
	switch s {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func getInfo(skip int) (fileFunc, filePath string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime err")
	}
	fileFunc = runtime.FuncForPC(pc).Name()
	fileFunc = strings.Split(fileFunc, ".")[1]
	filePath = path.Base(file)
	return
}
