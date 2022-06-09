package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       logLevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	errfileObj  *os.File
}

func NewFileLogger(loglevel, fp, fn string, maxSiez int64) *FileLogger {
	level, err := parseLogLevel(loglevel)
	if err != nil {
		panic("传入无效log等级")
	}
	f1 := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSiez,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *FileLogger) initFile() error {
	fullFileObj := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileObj, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("fileObj initFile in fullFileObj err", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileObj+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	f.fileObj = fileObj
	f.errfileObj = errfileObj
	return nil
}

func (f *FileLogger) enable(lv logLevel) bool {
	return lv >= f.Level
}
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("checkSize get info failed", err)
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	now := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("splitFile get info failed", err)
		return nil, err
	}
	fullName := path.Join(f.filePath, fileInfo.Name())
	newFileName := fmt.Sprintf("%sbak%s", fullName, now)
	file.Close()
	os.Rename(fullName, newFileName)
	fileObj, err := os.OpenFile(fullName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("create OpenFile err", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *FileLogger) log(lv logLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		s := getLogString(lv)
		nowStr := time.Now().Format("20060102150405")
		msg := fmt.Sprintf(format, a...)
		fileFunc, filePath, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			filelog, err := f.splitFile(f.fileObj)
			if err != nil {
				fmt.Println(" f.splitFile(f.fileObj) err", err)
			}
			f.fileObj = filelog
		}
		fmt.Fprintf(f.fileObj, "%s [%s %s %d] %s %s\n", s, fileFunc, filePath, lineNo, msg, nowStr)
		if lv > ERROR {
			if f.checkSize(f.fileObj) {
				errfilelog, err := f.splitFile(f.errfileObj)
				if err != nil {
					fmt.Println(" f.splitFile(f.fileObj) err", err)
				}
				f.errfileObj = errfilelog
			}
		}
		fmt.Fprintf(f.errfileObj, "%s [%s %s %d] %s %s\n", s, fileFunc, filePath, lineNo, msg, nowStr)
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
func (f *FileLogger) Close() {
	f.errfileObj.Close()
	f.fileObj.Close()
}
