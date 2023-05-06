package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 结构体模式 读取配置信息 可删除
var Conf = new(Config)

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Port int `yaml:"port"`
	Mode int `yaml:"mode"`
	Ver  int `yaml:"ver"`
}

type MySQLConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DB           string `yaml:"db"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	MaxConns     int    `yaml:"max_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type RedisConfig struct {
	DB   string `yaml:"db"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// 结构体模式 读取配置信息 可删除

// Init 当前为直接获取，结构体模式并存
func Init() (err error) {
	viper.SetConfigName("config")  //指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")    //指定配置文件类型
	viper.AddConfigPath(".")       //指定查找配置文件路径（这里使用相对路径）
	viper.AddConfigPath("./conf/") //指定查找配置文件路径（这里使用相对路径）
	err = viper.ReadInConfig()     //结构体模式 读取配置信息 可删除
	viper.Unmarshal(&Conf)         //结构体模式 读取配置信息 可删除
	viper.WatchConfig()            //监控配置文件更改
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件发生变更后会调用的回调函数
		viper.Unmarshal(&Conf) //结构体模式 读取配置信息 可删除 当配置文件变化后配置信息更新到全局变量Conf中
		fmt.Println("配置文件被修改:", e.Name)
	})
	return err // 不使用结构体模式替换为 return viper.ReadInConfig() //读取配置信息
}
