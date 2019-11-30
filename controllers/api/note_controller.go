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
	articles, cursor := services.NoteService.GetArticles(cursor)
	//return simple.JsonCursorData(render.BuildSimpleArticles(articles), strconv.FormatInt(cursor, 10))

	return simple.JsonData(articles)
}

// 文章详情
func (this *NoteController) GetBy(articleId int64) *simple.JsonResult {
	article := services.NoteService.Get(articleId)
	if article == nil || article.Status != model.ArticleStatusPublished {
		return simple.JsonErrorMsg("文章不存在")
	}
	return simple.JsonData(article)
}