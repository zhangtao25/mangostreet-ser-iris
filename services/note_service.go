package services

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
	"mangostreet-ser-iris/model"
	"mangostreet-ser-iris/repositories"
)

var NoteService = newNoteService()

func newNoteService() *noteService {
	return &noteService{

	}
}

type noteService struct {
}

func (this *noteService) Get(id int64) *model.Note {
	return repositories.NoteRepository.Get(simple.DB(), id)
}

// 发布文章
func (this *noteService) Publish(
	userId int64,
	title,
	content,
	contentType,
	images string) (note *model.Note, err error) {

	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)

	if len(title) == 0 {
		return nil, errors.New("标题不能为空")
	}

	note = &model.Note{
		UserId:      userId,
		Title:       title,
		Content:     content,
		ContentType: contentType,
		Status:      0,
		Images:      images,
		CreateTime:  simple.NowTimestamp(),
		UpdateTime:  simple.NowTimestamp(),
	}

	err = simple.Tx(simple.DB(), func(tx *gorm.DB) error {
		err := repositories.NoteRepository.Create(tx, note)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// 文章列表
func (this *noteService) GetArticles(cursor int64) (articles []model.Note, nextCursor int64) {
	cnd := simple.NewSqlCnd().Eq("status", model.ArticleStatusPublished).Desc("id").Limit(20)
	if cursor > 0 {
		cnd.Lt("id", cursor)
	}
	articles = repositories.NoteRepository.Find(simple.DB(), cnd)
	if len(articles) > 0 {
		nextCursor = articles[len(articles)-1].Id
	} else {
		nextCursor = cursor
	}
	return
}