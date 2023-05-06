package clear

import (
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"io/ioutil"
	"time"
)

type Delinfo struct {
	ReceiveTime string
	DeviceInfo  string // name N is capital
}

var infolist []Delinfo

const dstName = "./myFile_Device.txt"

func TimedClear() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	// 每3秒执行一次
	_, err := s.Every(7).Day().At("03:30").Do(ClearRecords)
	if err != nil {
		return
	}

	s.StartBlocking()
}

func ClearRecords() {
	suspend := infolist
	dataFromFile, _ := ioutil.ReadFile(dstName) // file name with extension .txt or .json
	now, err := time.Parse("2006-01-02 15-04", time.Now().Format("2006-01-02 15-04"))
	if err != nil {
		fmt.Println("时间转换失败")
		return
	}
	json.Unmarshal(dataFromFile, &infolist)
	for _, val := range infolist {
		historicaltime, err := time.Parse("2006-01-02 15-04", val.ReceiveTime)
		if err != nil {
			fmt.Println("时间转换失败")
			return
		}
		sub := now.Sub(historicaltime).Hours()
		fmt.Println(sub)
		if sub < 168 {
			fmt.Println("小于7天")
			suspend = append(suspend, Delinfo{
				DeviceInfo:  val.DeviceInfo,
				ReceiveTime: val.ReceiveTime,
			})
		}
	}

	finalData, _ := json.MarshalIndent(suspend, "", " ")
	ioutil.WriteFile(dstName, finalData, 0644)

	fmt.Println("deldata=", suspend)
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"strings"
//	"time"
//)
//
//type delinfo struct {
//	DeviceInfo string // name N is capital
//}
//
//var infolist []delinfo
//
//var deldata []delinfo
//
//func ClearInfo() {
//	suspend := infolist
//	data_from_file, _ := ioutil.ReadFile("myFile_Device.txt") // file name with extension .txt or .json
//	// unmarshall/parse data received from file and save/push in slice
//	// 2 argument 1. data source, 2. data slice to store data解组从文件接收的数据并在切片 2 参数中保存推送 1. 数据源，2. 用于存储数据的数据切片
//	now, err := time.Parse("2006-01-02 15-04", time.Now().Format("2006-01-02 15-04"))
//	if err != nil {
//		fmt.Println("时间转换失败")
//		return
//	}
//	json.Unmarshal(data_from_file, &infolist)
//	for _, vla := range infolist {
//		split := strings.Split(vla.DeviceInfo, ": = ")
//		fmt.Println(split[0])
//		historicaltime, err := time.Parse("2006-01-02 15-04", split[0])
//		if err != nil {
//			fmt.Println("时间转换失败")
//			return
//		}
//		sub := now.Sub(historicaltime).Hours()
//		fmt.Println(sub)
//		if sub < 168 {
//			fmt.Println("小于7天")
//			suspend = append(suspend, delinfo{
//				DeviceInfo: vla.DeviceInfo,
//			})
//		}
//	}
//	fmt.Println(infolist)
//	fmt.Println("deldata=", suspend)
//}
