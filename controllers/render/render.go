package render

import (
	//"html/template"
	//"strconv"
	"strings"

	//"github.com/tidwall/gjson"

	//"github.com/PuerkitoBio/goquery"
	//"github.com/mlogclub/simple"

	//"mangostreet-ser-iris/common"
	"mangostreet-ser-iris/common/avatar"
	//"mangostreet-ser-iris/common/urls"
	"mangostreet-ser-iris/model"
	//"mangostreet-ser-iris/services"
	//"mangostreet-ser-iris/services/cache"
)


func BuildUser(user *model.User) *model.UserInfo {
	if user == nil {
		return nil
	}
	a := user.Avatar
	if len(a) == 0 {
		a = avatar.DefaultAvatar
	}
	roles := strings.Split(user.Roles, ",")
	return &model.UserInfo{
		Id:          user.Id,
		Username:    user.Username.String,
		Nickname:    user.Nickname,
		Avatar:      a,
		Email:       user.Email.String,
		Type:        user.Type,
		Roles:       roles,
		Description: user.Description,
		PasswordSet: len(user.Password) > 0,
		CreateTime:  user.CreateTime,
	}
}