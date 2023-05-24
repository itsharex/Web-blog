package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

type AskBoxService struct {
	askboxDao *dao.AskBoxDao
}

func NewAskBoxService() *AskBoxService {
	return &AskBoxService{
		askboxDao: dao.NewAskBoxDao(),
	}
}

// 报错？
//func (a *AskBoxService) GetAnsweredQA([]model.Askbox, error) {
//	return a.askboxDao.GetAnsweredQA()
//}

func (a *AskBoxService) AddNewQuestion(askbox *model.Askbox) error {
	return a.askboxDao.AddNewQuestion(askbox)
}

func (a *AskBoxService) AppendOldQuestion(askbox *model.Askbox) error {
	return a.askboxDao.AddNewQuestion(askbox)
}

func (a *AskBoxService) ClickLikes(likes int, parentID int, childID int) error {
	return a.askboxDao.ClickLikes(likes, parentID, childID)
}
