package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Context 不想再有什么变化就使用struat
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmedPassword"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (c Context) ReadJson(req interface{}) error {
	//读取body
	//序列化
	r := c.R
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, req)
	if err != nil {
		return err
	}
	return nil
}

func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}
	resp := &commonResponse{Data: 123}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest, data)
}
func (c *Context) SystemErrJson(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, data)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer,
		R: request,
	}
}
