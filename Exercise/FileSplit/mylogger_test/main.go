package main

import (
	"github.com/v04007/Exercise/mylogger"
)

func main() {
	log := mylogger.NewFileLogger("debug", "./", "case", 10*1024)
	for {
		age := 100
		name := "理想"
		log.Debug("这是debug等级年龄：%d,名字：%s", age, name)
		log.Info("这是Info等级")
		log.Warning("这是Warning等级")
		log.Error("这是Error等级")
		log.Fatal("这是Fatal等级")
	}

}
