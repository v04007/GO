package mylogger

// import (
// 	"fmt"
// 	"time"
// )

// type LogLevel uint64

// type ConsoleLogger struct {
// 	Level LogLevel
// }

// func Newlog(levelStr string) ConsoleLogger {
// 	level, err := parseLogLevel(levelStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ConsoleLogger{
// 		Level: level,
// 	}
// }

// func (c ConsoleLogger) enable(logLevel LogLevel) bool {
// 	return logLevel >= c.Level
// }

// func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
// 	if c.enable(lv) {
// 		msg := fmt.Sprintf(format, a...)
// 		now := time.Now().Format("2006-01-02-15 04:05:00")
// 		fucnName, fileName, lineNo := getInfo(3)
// 		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, fucnName, lineNo, msg)
// 	}
// }

// func (c ConsoleLogger) Debug(format string, a ...interface{}) {
// 	c.log(DEBUG, format, a...)
// }
// func (c ConsoleLogger) Info(format string, a ...interface{}) {
// 	c.log(INFO, format, a...)
// }
// func (c ConsoleLogger) Warning(format string, a ...interface{}) {
// 	c.log(WARNING, format, a...)
// }
// func (c ConsoleLogger) Error(format string, a ...interface{}) {
// 	c.log(ERROR, format, a...)
// }
// func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
// 	c.log(FATAL, format, a...)
// }
