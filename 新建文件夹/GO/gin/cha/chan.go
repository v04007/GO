package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name    string
	Message string
	Number  int
}

func main() {
	r := gin.Default()
	r.GET("/chan", func(c *gin.Context) {
		query := c.Query("data")
		c.JSON(http.StatusOK, gin.H{"type": "JSON", "status": 200, "data": query})
	})
	r.Run("localhost:8000")
}
USE [master]
EXEC sp_configure 'show advanced option', '1' --只有这个高级选项被打开的时候，才有权限修改其他配置。
go
RECONFIGURE     --运行RECONFIGURE语句进行安装，也就是说，使以上语句生效
go
EXEC sp_configure 'Ole Automation Procedures';
GO

--启用Ole Automation Procedures
sp_configure 'show advanced options', 1;
GO
RECONFIGURE;
GO
sp_configure 'Ole Automation Procedures', 1;
GO
RECONFIGURE;
GO


ALTER trigger [dbo].[name]
on [dbo].[TBSF_Name]
after insert
as
begin
declare @ServiceUrl as varchar(1000)
if exists(select tbsf_name.id from tbsf_name, inserted
where tbsf_name.id = inserted.id)
begin
set @ServiceUrl = 'http://localhost:8000?data=' + '哈哈哈' + '&message='

Declare @Object as Int
Declare @ResponseText as Varchar(8000)
--通过http协议调用的接口地址'
Exec sp_OACreate 'MSXML2.XMLHTTP', @Object OUT;
Exec sp_OAMethod @Object, 'open', NULL, 'get',@ServiceUrl, 'false'
Exec sp_OAMethod @Object, 'send'
Exec sp_OAMethod @Object, 'responseText', @ResponseText OUTPUT

update tbsf_name set id = '-10'
from tbsf_name, inserted
end
end

select * from TBSF_Name