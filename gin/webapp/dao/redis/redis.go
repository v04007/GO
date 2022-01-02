package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"webapp/setting"
)

var (
	Client *redis.Client
)

// Init 初始化连接
func Init() (err error) {
	Client = redis.NewClient(&redis.Options{
		//Addr: fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),//直接获取值
		Addr: fmt.Sprintf("%s:%d", setting.Conf.RedisConfig.Host, setting.Conf.RedisConfig.Port), //结构体模式，不使用可删除
		DB:   viper.GetInt("redis.db"),                                                           //use default DB
	})
	_, err = Client.Ping().Result()
	return
}
