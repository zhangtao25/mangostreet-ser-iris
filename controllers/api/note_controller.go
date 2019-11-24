package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"

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

	//userId int64,
	//	title,
	//	content,
	//	contentType,
	//	images string

	note, err := services.NoteService.Publish(user.Id, title, content, contentType, images,)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(note)
}