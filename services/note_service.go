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

	//if err == nil {
	//	baiduseo.PushUrl(urls.ArticleUrl(article.Id))
	//}
	return
}