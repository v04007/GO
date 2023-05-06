package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/dao/mysql"
	"webapp/logger"
	"webapp/pkg/snowflake"
	"webapp/routers"
	"webapp/setting"
)

func main() {
	//加载配置信息
	if err := setting.Init(); err != nil {
		fmt.Printf("load setting failed,err:%#v\n", err)
		return
	}
	//初始化日志库
	logger.Init()
	defer zap.L().Sync() //将缓冲区log日志落盘
	zap.L().Info("load setting success")
	//初始化mysql连接
	if err := mysql.Init(); err != nil {
		zap.L().Error("init mysql failed,err%v\n",
			zap.Error(err),
			zap.String("mysql", "xxxx"),
			zap.Int("port", 3306))
		return
	}
	defer mysql.Close()

	//zap.L().Info("init mysql success")
	////初始化redis连接
	//if err := redis.Init(); err != nil {
	//	zap.L().Error("init redis failed", zap.Error(err))
	//	return
	//}
	//redis.Client.Close()
	//zap.L().Info("init redis success")

	//初始化ID生成器
	if err := snowflake.Init(uint16(viper.GetInt("app.machine_id"))); err != nil {
		zap.L().Error("init snowflake failed", zap.Error(err))
	}
	//kafka.init
	//etcd.init
	//加载路由信息
	r := routers.SetupRouters()
	//err := r.Run(fmt.Sprintf("localhost:%d", viper.GetInt("app.port")))

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", viper.GetInt("app.port")),
		Handler: r,
	}
	//开启一个goroutine 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("%s\n", zap.Error(err))
		}
	}()

	//等待终端信号来优雅地关闭服务器，为关闭服务其操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) //创建一个接受信号的通道
	//kilL默认会发送syscall.SIGTERM信号
	//kill -2发送syscall.SIGINT信号,我们常用的CtrL+C就是敏发系统SIGINT信号
	//kilL -9发送syscall.SIGKILL信号,但是不能被補获,所以不需要添加它
	//signal.Notify把收到的syscall.SIGINT 或syscall.SIGTERM信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //此处不会阻塞
	<-quit                                               //阻塞在此，但接受到上述两种信号时才会继续往执行
	zap.L().Info("Shutdown Server...")
	//创建一个5秒超时的context
	//相当于告诉程序我给你5秒钟的时间你把没完成的请求处理一下,之后就要关机了
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//5秒优雅关闭服务器（将未处理外的请求处理完再关闭服务其），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("server exiting")
}
