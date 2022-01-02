package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//var (
//	logger *zap.Logger
//)

//Init 初始化日志
func Init() {
	writeSyncer := getLogWriter() //日志输出位置相关
	errWriteSyncer := getErrLogWriter()
	encoder := getEncoder()    //日志输入格式相关
	level := zap.AtomicLevel{} //解析字符串格式级别
	if err := level.UnmarshalText([]byte(viper.GetString("log.level"))); err != nil {
		level = zap.NewAtomicLevel() //默认使用info
	}
	//core := zapcore.NewCore(encoder, writeSyncer, l)

	//根据app模式把日志输出到不同位置
	var core zapcore.Core
	if viper.GetString("app.mode") == gin.DebugMode { //debug模式
		//consoleEncoder 一个往终端输出的配置对象
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		//NewTee 可以指定多个日志配置
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			//创建一个将debug级别往上的日志输入,输出到终端配置信息
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			//将error级别以上日志输出到err文件
			zapcore.NewCore(encoder, errWriteSyncer, zapcore.ErrorLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}
	//logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) //如果在原方法上封装一层此处要加这一句
	//zap.Fields(zap.String("app", "xiaodebu"))不同业务线有不同业务线区分
	logger := zap.New(core, zap.AddCaller(), zap.Fields(zap.String("app", "xiaodebu"))) //根据上面配置创建logger,zap.AddCaller()那个函数调用
	zap.ReplaceGlobals(logger)                                                          //替换zap库里面全局的logger
	//sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp" //将ts修改为自定义
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //人们可读时间格式
	//默认格式为1625829519.1739724 float64(nanos),需要自己修改为int类型,如下
	//sec := int64(nanos) / int64(time.Second)
	//enc.AppendInt64(sec)
	encoderConfig.EncodeTime = zapcore.EpochTimeEncoder //时间戳
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig) //刻度日志
	return zapcore.NewJSONEncoder(encoderConfig) //json格式日志
}

//使用zapcore包创建
//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.Create("./test.log")
//	return zapcore.AddSync(file)
//}

//使用第三方库实现日志切割创建日志库
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.max_size"),    //日志文件大小 单位:MB
		MaxBackups: viper.GetInt("log.max_backups"), //备份数量
		MaxAge:     viper.GetInt("log.max_age"),     //备份时间 单位: 天
		Compress:   viper.GetBool("log.compre"),     //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getErrLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("log.filename") + "err",
		MaxSize:    viper.GetInt("log.max_size"),    //日志文件大小 单位:MB
		MaxBackups: viper.GetInt("log.max_backups"), //备份数量
		MaxAge:     viper.GetInt("log.max_age"),     //备份时间 单位: 天
		Compress:   viper.GetBool("log.compre"),     //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
