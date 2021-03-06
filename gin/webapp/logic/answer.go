package logic

import (
	"errors"
	"go.uber.org/zap"
	"html"
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/snowflake"
)

var (
	ErrorInvalidQuestionID = errors.New("无效的问题ID")
)

func CreateAnswer(answer *models.Answer) (err error) {
	//1. 针对content做一个转义,防止xss漏洞
	answer.Content = html.EscapeString(answer.Content)
	//敏感词过滤
	//2.校验questionID是否存在
	if ok := mysql.CheckQuestionIDExist(answer.QuestionID); !ok {
		err = ErrorInvalidQuestionID
		zap.L().Warn("invalid question id", zap.Uint64("questionID", answer.QuestionID))
		return
	}

	//3.生成回答的id
	answerID, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	answer.AnswerID = answerID
	err = mysql.CreateAnswer(answer)
	if err != nil {
		zap.L().Error("snowflake.CreateAnswer failed", zap.Error(err))
		return
	}
	return
}

func GetAnswerList(questionID uint64, offset, limit int) (apiAnswer *models.ApiAnswerList, err error) {
	answerList, err := mysql.GetAnswerList(questionID, offset, limit)
	if err != nil {
		zap.L().Error("mysql.GetAnswerList failed", zap.Error(err))
		return
	}
	if len(answerList) == 0 {
		return
	}

	var userIDList []uint64
	for _, v := range answerList {
		userIDList = append(userIDList, v.AuthorID)
	}

	userInfoList, err := mysql.GetUserInfoList(userIDList)
	if err != nil {
		zap.L().Error("mysql.GetUserInfoList failed",
			zap.Any("userIDList", userIDList),
			zap.Error(err),
		)
		return
	}

	apiAnswerList := &models.ApiAnswerList{
		AnswerList: make([]*models.ApiAnswer, 0, len(answerList)),
	}
	for _, v := range answerList {
		apiAnswer := &models.ApiAnswer{}
		apiAnswer.Answer = *v

		for _, user := range userInfoList {
			if user.UserID == v.AuthorID {
				apiAnswer.AuthorName = user.UserName
				break
			}
		}
		apiAnswerList.AnswerList = append(apiAnswerList.AnswerList, apiAnswer)
	}
	count, err := mysql.GetAnswerCount(questionID)
	if err != nil {
		zap.L().Error("mysql.GetQuestionCount failed",
			zap.Uint64("questionID", questionID),
			zap.Error(err),
		)
		return
	}
	apiAnswerList.TotalCount = count
	return
}
