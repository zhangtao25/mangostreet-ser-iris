package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"mangostreet-ser-iris/model"
	"mangostreet-ser-iris/services"
)

type NoteController struct {
	Ctx iris.Context
}


// 发表文章
func (this *NoteController) PostCreate() *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(this.Ctx)
	if user == nil {
		return simple.JsonError(simple.ErrorNotLogin)
	}
	var (
		title   = this.Ctx.PostValue("title")
		content = this.Ctx.PostValue("content")
		contentType = this.Ctx.PostValue("contentType")
		images = this.Ctx.PostValue("images")
	)

	note, err := services.NoteService.Publish(user.Id, title, content, contentType, images,)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(note)
}

// 文章列表
func (this *NoteController) GetNotes() *simple.JsonResult {
	cursor := simple.FormValueInt64Default(this.Ctx, "cursor", 0)
	_, cursor, noteDetail := services.NoteService.GetArticles(cursor)

	return simple.JsonData(noteDetail)
}

// 文章详情
func (this *NoteController) GetBy(noteId int64) *simple.JsonResult {
	note := services.NoteService.Get(noteId)
	uid := note.UserId
	user := services.UserService.Get(uid)


	if note == nil || note.Status != model.ArticleStatusPublished {
		return simple.JsonErrorMsg("文章不存在")
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
	var noteDetail NoteDetail
	noteDetail.Id          = note.Id
	noteDetail.Title       = note.Title
	noteDetail.Content     = note.Content
	noteDetail.ContentType = note.ContentType
	noteDetail.Status      = note.Status
	noteDetail.Images      = note.Images
	noteDetail.CreateTime  = note.CreateTime
	noteDetail.UpdateTime  = note.UpdateTime
	noteDetail.Nickname    = user.Nickname
	noteDetail.Avatar      = user.Avatar

	return simple.JsonData(noteDetail)
}