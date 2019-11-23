package model

import "database/sql"

var Models = []interface{}{
	&User{}, &UserToken{}, &Category{},
	// &Tag{}, &Article{}, &ArticleTag{}, &Comment{}, &Favorite{},
	//&Topic{}, &TopicTag{}, &TopicLike{}, &Message{}, &SysConfig{}, &Project{}, &Subject{}, &SubjectContent{}, &Link{},
	//&CollectRule{}, &CollectArticle{}, &ThirdAccount{},
}

type Model struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

const (
	UserStatusOk       = 0
	UserStatusDisabled = 1

	UserTokenStatusOk       = 0
	UserTokenStatusDisabled = 1

	UserTypeNormal = 0 // 普通用户
	UserTypeGzh    = 1 // 公众号用户

	CategoryStatusOk       = 0
	CategoryStatusDisabled = 1

	TagStatusOk       = 0
	TagStatusDisabled = 1

	ArticleStatusPublished = 0 // 已发布
	ArticleStatusDeleted   = 1 // 已删除
	ArticleStatusDraft     = 2 // 草稿

	ArticleTagStatusOk      = 0
	ArticleTagStatusDeleted = 1

	TopicStatusOk      = 0
	TopicStatusDeleted = 1

	TopicTagStatusOk      = 0
	TopicTagStatusDeleted = 1

	ContentTypeHtml     = "html"
	ContentTypeMarkdown = "markdown"

	CommentStatusOk      = 0
	CommentStatusDeleted = 1

	EntityTypeArticle = "article"
	EntityTypeTopic   = "topic"

	MsgStatusUnread = 0 // 消息未读
	MsgStatusReaded = 1 // 消息已读

	MsgTypeComment = 0 // 回复消息

	LinkStatusOk      = 0 // 正常
	LinkStatusDeleted = 1 // 删除
	LinkStatusPending = 2 // 待审核

	CollectRuleStatusOk       = 0 // 启用
	CollectRuleStatusDisabled = 1 // 禁用

	CollectArticleStatusPending   = 0 // 待审核
	CollectArticleStatusAuditPass = 1 // 审核通过
	CollectArticleStatusAuditFail = 2 // 审核失败
	CollectArticleStatusPublished = 3 // 已发布

	ThirdAccountTypeGithub = "github"
	ThirdAccountTypeQQ     = "qq"
)

type User struct {
	Model
	Username    sql.NullString `gorm:"size:32;unique;" json:"username" form:"username"`
	Email       sql.NullString `gorm:"size:128;unique;" json:"email" form:"email"`
	Nickname    string         `gorm:"size:16;" json:"nickname" form:"nickname"`
	Avatar      string         `gorm:"type:text" json:"avatar" form:"avatar"`
	Password    string         `gorm:"size:512" json:"password" form:"password"`
	Status      int            `gorm:"index:idx_status;not null" json:"status" form:"status"`
	Roles       string         `gorm:"type:text" json:"roles" form:"roles"`
	Type        int            `gorm:"not null" json:"type" form:"type"`
	Description string         `gorm:"type:text" json:"description" form:"description"`
	CreateTime  int64          `json:"createTime" form:"createTime"`
	UpdateTime  int64          `json:"updateTime" form:"updateTime"`
}

type UserToken struct {
	Model
	Token      string `gorm:"size:32;unique;not null" json:"token" form:"token"`
	UserId     int64  `gorm:"not null;index:idx_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"not null;index:idx_status" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
}

// 分类
type Category struct {
	Model
	Name        string `gorm:"size:32;unique;not null" json:"name" form:"name"`
	Description string `gorm:"size:1024" json:"description" form:"description"`
	Status      int    `gorm:"index:idx_status;not null" json:"status" form:"status"`
	CreateTime  int64  `json:"createTime" form:"createTime"`
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`
}