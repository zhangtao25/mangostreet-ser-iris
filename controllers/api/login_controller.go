package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"

	"mangostreet-ser-iris/controllers/render"
	"mangostreet-ser-iris/model"
	"mangostreet-ser-iris/services"
)

type LoginController struct {
	Ctx iris.Context
}

// 注册
func (this *LoginController) PostSignup() *simple.JsonResult {
	var (
		username   = this.Ctx.PostValueTrim("username")
		password   = this.Ctx.PostValueTrim("password")
		rePassword = this.Ctx.PostValueTrim("rePassword")
		nickname   = this.Ctx.PostValueTrim("nickname")
		ref        = this.Ctx.FormValue("ref")
	)
	user, err := services.UserService.SignUp(username, "", nickname, password, rePassword)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return this.GenerateLoginResult(user, ref)
}

// 用户名密码登录
func (this *LoginController) PostSignin() *simple.JsonResult {
	var (
		username = this.Ctx.PostValueTrim("username")
		password = this.Ctx.PostValueTrim("password")
		ref      = this.Ctx.FormValue("ref")
	)
	user, err := services.UserService.SignIn(username, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return this.GenerateLoginResult(user, ref)
}

// 退出登录
func (this *LoginController) GetSignout() *simple.JsonResult {
	err := services.UserTokenService.Signout(this.Ctx)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}

// user: login user, ref: 登录来源地址，需要控制登录成功之后跳转到该地址
func (this *LoginController) GenerateLoginResult(user *model.User, ref string) *simple.JsonResult {
	token, err := services.UserTokenService.Generate(user.Id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.NewEmptyRspBuilder().
		Put("token", token).
		Put("user", render.BuildUser(user)).
		Put("ref", ref).JsonResult()
}
