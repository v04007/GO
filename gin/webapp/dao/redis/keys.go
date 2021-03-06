package redis

const (
	AppPrefix = "xiaodebu:"

	KeyQuestionTimeZSet  = AppPrefix + "question:time:"  //发布时间为分数的问题ZSet
	KeyQuestionScoreZSet = AppPrefix + "question:score:" //投票累计作为分数的问题ZSet

	KeyQuestionVotedSetPrefix    = AppPrefix + "question:voted:" //某个问题已经投票的用户set
	KeyQuestionInfoHashPrefix    = AppPrefix + "question:info:"  //hash,存储问题基础信息
	KeyCategoryQuestionSetPrefix = AppPrefix + "question:category:"
)
