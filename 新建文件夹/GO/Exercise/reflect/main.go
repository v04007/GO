package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"address"`
	Port     int    `ini:"port"`
	Password string `ini:"username"`
	Database string `ini:"password"`
	Test     bool   `ini:"test"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

// 取类型信息typeof取值信息用valueof
func loadini(fileName string, data interface{}) (err error) {
	// 传入参数应该为指针类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	// 传入的参数必须为结构体类型指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}
	// 读文件的到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 将字节类型转换为字符串
	lineSlice := strings.Split(string(b), "\r\n")
	// fmt.Printf("%#v", lineSlice)
	// 一行一行读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾空格
		line := strings.TrimSpace(line)
		// 如果为空就跳过
		if len(line) == 0 {
			continue
		}
		// 如果是注释就去掉,HasPrefix方法以什么方式开头
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue //跳出循环
		}
		// 如果是[开头就代表是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行首尾[]去掉，渠道中间的内容把首尾空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根基sectionName去data里面根据反射找到对应结构体

			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应结构体，把字段记录下来
					structName = field.Name
					fmt.Printf("找到%s对应嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 如果不是[开头就是=分割键值对
			// 按照=分割这一行，等号左边是key右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//根据structName去data里把对应嵌套结构图取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //根据给定字符串返回字符串对应的结构体字段的信息
			sType := sValue.Type()                     //Type 返回 v 的类型。

			if sType.Kind() != reflect.Struct { //值的种类
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 遍历嵌套结构体每一个字段判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) //根据索引，返回索引对应的结构体字段的信息。
				fileType = filed
				if filed.Tag.Get("ini") == key {
					fieldName = filed.Name
					break
				}
			}
			// 如果key=tag，给这个字段赋值
			// 根据fieldName去去除这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应字符
				continue
			}
			fileObj := sValue.FieldByName(fieldName) //赋值

			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}

		}
	}
	return nil
}

func main() {
	var cfg Config
	err := loadini("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
	}
	fmt.Printf("%#v", cfg)
}
