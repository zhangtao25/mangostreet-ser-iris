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

type NoteDetail struct {
	Id          int64       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
	Title       string		`gorm:"size:128;not null;" json:"title" form:"title"`
	Content     string		`gorm:"type:longtext;not null;" json:"content" form:"content"`
	ContentType string		`gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`
	Status      int			`gorm:"int;not null" json:"status" form:"status"`
	Images   	string		`gorm:"type:text" json:"images" form:"images"`
	CreateTime  int64		`json:"createTime" form:"createTime"`
	UpdateTime  int64		`json:"updateTime" form:"updateTime"`
	Nickname    string		`gorm:"size:16;" json:"nickname" form:"nickname"`
	Avatar      string		`gorm:"type:text" json:"avatar" form:"avatar"`
}

// 文章列表
func (this *noteService) GetArticles(cursor int64) (articles []model.Note, nextCursor int64, noteDetails []NoteDetail) {
	cnd := simple.NewSqlCnd().Eq("status", model.ArticleStatusPublished).Desc("id").Limit(20)
	if cursor > 0 {
		cnd.Lt("id", cursor)
	}
	articles = repositories.NoteRepository.Find(simple.DB(), cnd)



	//var noteDetails []NoteDetail

	for i := 0; i < len(articles); i++ {

		user := repositories.UserRepository.Get(simple.DB(), articles[i].UserId)

		var noteDetail NoteDetail
		noteDetail.Id = articles[i].Id

		noteDetail.Id          = articles[i].Id
		noteDetail.Title       = articles[i].Title
		noteDetail.Content     = articles[i].Content
		noteDetail.ContentType = articles[i].ContentType
		noteDetail.Status      = articles[i].Status
		noteDetail.Images      = articles[i].Images
		noteDetail.CreateTime  = articles[i].CreateTime
		noteDetail.UpdateTime  = articles[i].UpdateTime
		noteDetail.Nickname    = user.Nickname
		noteDetail.Avatar      = user.Avatar

		noteDetails = append(noteDetails, noteDetail)
	}

	if len(articles) > 0 {
		nextCursor = articles[len(articles)-1].Id
	} else {
		nextCursor = cursor
	}
	return
}