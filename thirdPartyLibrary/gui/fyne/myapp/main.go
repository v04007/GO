package main

import (
	"MODULE_NAME/clear"
	"MODULE_NAME/logger"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/App"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type ReceiveData struct {
	ReceiveTime string
	DeviceInfo  string // name N is capital
}

//type StudentList []Student

var (
	datasheets []ReceiveData
)

const (
	//addr = "10.28.0.51:4003"
	addr     = "127.0.0.1:6390"
	sendAddr = "http://10.28.83.123:8081/v1/adm"
)

func main() {
	logger.Init()
	defer zap.L().Sync()
	go clear.TimedClear()                                     //定时每7天清除日志
	data_from_file, _ := ioutil.ReadFile("myFile_Device.txt") // file name with extension .txt or .json
	// unmarshall/parse data received from file and save/push in slice
	// 2 argument 1. data source, 2. data slice to store data解组从文件接收的数据并在切片 2 参数中保存推送 1. 数据源，2. 用于存储数据的数据切片
	json.Unmarshal(data_from_file, &datasheets)
	a := app.New()
	// new title and window
	w := a.NewWindow("CRUD APP")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	hello := widget.NewLabel("Hello!")
	content := make(chan ReceiveData)
	//hello := widget.NewLabel("Hello Fyne!")
	go func() {
		for {
			time.Sleep(time.Second * 1)
			select {
			case data := <-content:
				if len(data.ReceiveTime) != 0 && len(data.DeviceInfo) != 0 {
					myData := &ReceiveData{
						ReceiveTime: data.ReceiveTime,
						DeviceInfo:  data.DeviceInfo,
					}
					datasheets = append(datasheets, *myData)
					final_data, _ := json.MarshalIndent(datasheets, "", " ")
					ioutil.WriteFile("myFile_Device.txt", final_data, 0644)
				}

			default:

			}
		}
	}()
	fmt.Println("程序开始2")
	scroll := container.Scroll{}
	scroll.CreateRenderer()
	//clock := widget.NewLabel(str)
	//text := canvas.NewText("str", color.Black)
	//text.Alignment = fyne.TextAlignTrailing
	//text.TextStyle = fyne.TextStyle{Italic: true}

	//var list *widget.List

	final_data, _ := json.MarshalIndent(datasheets, "", " ")
	ioutil.WriteFile("myFile_Device.txt", final_data, 0644)

	list := widget.NewList(
		// first argument is item count
		// len() function to get myStudentData slice len第一个参数是项目计数 len() 函数，用于获取 myStudentData 切片 len
		func() int { return len(datasheets) },
		// 2nd argument is for widget choice. I want to use label第二个参数用于小部件选择。我想使用标签
		func() fyne.CanvasObject { return widget.NewLabel("") },
		// 3rd argument is to update widget with our new data第三个参数是用我们的新数据更新小部件
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(fmt.Sprintf("%s:%s", datasheets[lii].ReceiveTime, datasheets[lii].DeviceInfo))
		},
	)

	startbtn := widget.NewButton("Click to start", func() {
		go tcpServer(content)
		hello.SetText("start listening")
	})

	//startbtn.MinSize()
	//w.SetContent(
	//	// lets create Hsplite container让我们创建 Hsplite 容器
	//	container.NewHSplit(
	//		// first argument is list of data第一个参数是数据列表
	//		list,
	//		// 2nd is
	//		// vbox container
	//		container.NewVBox(startbtn),
	//	),
	//)

	c := container.NewVScroll(
		container.NewVSplit( // New Horizontal Box
			container.NewVBox(startbtn, hello),
			list,
		))

	c.Direction = container.ScrollHorizontalOnly
	// setup content
	w.SetContent(c)

	w.ShowAndRun()
}

func tcpServer(connop chan ReceiveData) {
	const (
		ip   = "127.0.0.1"
		port = 4003
	)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		//fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go handleConn(conn, connop)
	}
}
func handleConn(conn net.Conn, val chan ReceiveData) {
	reader := bufio.NewReader(conn)
	for {
		dat, err := reader.ReadSlice('&')
		if err == io.EOF {
			fmt.Println("读取完毕。。。。。。")
			break
		}
		val <- ReceiveData{
			ReceiveTime: time.Now().Format("2006-01-02 15-04"),
			DeviceInfo:  string(dat),
		}
		Send(string(dat), sendAddr)
		conn.Write([]byte("end"))
		//resultmap := internal.StringHandling(string(dat))
		//for key, val := range resultmap {
		//	internal.Send(key, val, sendAddr)
		//}
		//conn.Write([]byte("end"))
	}
}

func Send(data, sendAddr string) {
	body, err := json.Marshal(data)
	if err != nil {
		zap.L().Error("json序列化失败")
		println("序列化失败")
		return
	}

	resp, err := http.Post(sendAddr, "application/json", bytes.NewBuffer(body))
	if err != nil {
		zap.L().Error("发送失败", zap.Error(err))
		fmt.Println("发送失败")
		return
	}
	//读取响应后，需要关闭响应流。因此推迟关闭。它会自动处理它。
	defer resp.Body.Close()

	//Check response code, if New user is created then read response.
	switch resp.StatusCode {
	case http.StatusOK:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			zap.L().Error("读取响应失败", zap.Error(err))
			//Failed to read response.
			panic(err)
		}
		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)
	case http.StatusBadRequest:
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)
		zap.L().Error("写入失败400", zap.Error(err))
		fmt.Println("Response: ", resp.StatusCode)
	default:
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)
		fmt.Println("未知状态", resp.Status)
	}
}
