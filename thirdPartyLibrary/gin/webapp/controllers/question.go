package controllers

/*多数校验-->校验必要的数据是否存在
防XSS攻击 --> 'htmL/template"内容检测和转义
敏感词过滤-->涉黄、涉恐等等屏蔽问
数据库创逮记录-->dao层创建记录
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"webapp/logic"
	"webapp/models"
)

// QuestionSubmitHandler 提交问题
func QuestionSubmitHandler(c *gin.Context) {
	var question = new(models.Question)
	err := c.BindJSON(question)
	fmt.Println("参数:", question)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}

	if len(question.Caption) == 0 || len(question.Content) == 0 {
		ResponseError(c, CodeInvalidParams)
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("user is not login")
		ResponseError(c, CodeNotLogin)
		return
	}
	question.AuthorID = userID

	if err := logic.CreateQuestion(question); err != nil {
		zap.L().Error("create question failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	zap.L().Debug("create question success", zap.Any("question", question))
	ResponseSuccess(c, question.QuestionID)
}

// QuestionListsHandler 查看文体
func QuestionListsHandler(c *gin.Context) {
	offset, limit := getOffsetAndLimit(c)
	zap.L().Debug("get question list param success",
		zap.Int("offset", offset),
		zap.Int("limit", limit),
	)
	data, err := logic.GetQuestionList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetQuestionList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println(data, offset, limit)
	zap.L().Debug("get question list success", zap.Any("data", data))
	ResponseSuccess(c, data)
}

func QuestionDetailHandler(c *gin.Context) {
	questionID, err := getIDFromQuery(c, "question_id")
	if err != nil {
		zap.L().Error("invalid query param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetQuestionDetail(questionID)
	if err != nil {
		zap.L().Error("get question detail failed", zap.Error(err))
	}
	ResponseSuccess(c, data)
}
