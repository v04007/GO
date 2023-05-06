package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Init() {
	writeSyncer := getLogWriter()
	errWriteSyncer := getErrLogWriter()
	encoder := getEncoder()    //日志输入格式相关
	level := zap.AtomicLevel{} //解析字符串格式级别
	if err := level.UnmarshalText([]byte("info")); err != nil {
		level = zap.NewAtomicLevel() //默认使用info
	}
	var core zapcore.Core

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core = zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, level),
		//创建一个将debug级别往上的日志输入,输出到终端配置信息
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		//将error级别以上日志输出到err文件
		zapcore.NewCore(encoder, errWriteSyncer, zapcore.ErrorLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.Fields(zap.String("app", "adm"))) //根据上面配置创建logger,zap.AddCaller()那个函数调用
	zap.ReplaceGlobals(logger)                                                     //替换zap库里面全局的logger

}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./log/%s.log", "AoiDeviceDataInfo"),
		MaxSize:    10,    //日志的最大大小（M）
		MaxBackups: 5,     //日志的最大保存数量
		MaxAge:     30,    //日志文件存储最大天数
		Compress:   false, //是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getErrLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./log/%s.log", "AoiDeviceDataInfo") + "err",
		MaxSize:    10,    //日志的最大大小（M）
		MaxBackups: 5,     //日志的最大保存数量
		MaxAge:     30,    //日志文件存储最大天数
		Compress:   false, //是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"                 //将ts修改为自定义
	encoderConfig.EncodeTime = zapcore.EpochTimeEncoder //时间戳
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig) //json格式日志
}
