package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	"mangostreet-ser-iris/common/oss"
	"mangostreet-ser-iris/common/request"
	"mangostreet-ser-iris/services"
)

const uploadMaxBytes int64 = 1024 * 1024 * 3 // 1M

type UploadController struct {
	Ctx iris.Context
}

func (this *UploadController) Post() *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(this.Ctx)
	if user == nil {
		return simple.JsonError(simple.ErrorNotLogin)
	}

	//这下面是我自己写的QWQ
	var upload_path = "./upload/note/"
	//获取文件内容 要这样获取
	file, head, err := this.Ctx.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return simple.JsonErrorMsg(err.Error())
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create(upload_path + head.Filename)
	if err != nil {
		fmt.Println("文件创建失败")
		return simple.JsonErrorMsg(err.Error())
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return simple.JsonErrorMsg(err.Error())
	}
	err = request.PostFile(upload_path+head.Filename, "http://mangostreet.top:8001/upload/note")
	//删除文件
	//Go中删除文件和删除文件夹同一个函数
	//errRemove := os.Remove(upload_path + head.Filename);
	//if errRemove != nil {
	//	log.Fatal(errRemove);
	//}
	//if err !=nil {
	//	return simple.JsonErrorMsg(err.Error())
	//}
	return simple.NewEmptyRspBuilder().Put("urls", head.Filename).JsonResult()
}

// vditor上传
func (this *UploadController) PostEditor() {
	errFiles := make([]string, 0)
	succMap := make(map[string]string)

	user := services.UserTokenService.GetCurrent(this.Ctx)
	if user == nil {
		_, _ = this.Ctx.JSON(iris.Map{
			"msg":  "请先登录",
			"code": 1,
			"data": iris.Map{
				"errFiles": errFiles,
				"succMap":  succMap,
			},
		})
		return
	}

	maxSize := this.Ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := this.Ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = this.Ctx.WriteString(err.Error())
		return
	}

	form := this.Ctx.Request().MultipartForm
	files := form.File["file[]"]
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			logrus.Error(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}
		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			logrus.Error(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}
		url, err := oss.PutImage(fileBytes)
		if err != nil {
			logrus.Error(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}

		succMap[file.Filename] = url
	}

	_, _ = this.Ctx.JSON(iris.Map{
		"msg":  "",
		"code": 0,
		"data": iris.Map{
			"errFiles": errFiles,
			"succMap":  succMap,
		},
	})
	return

}

// vditor 拷贝第三方图片
func (this *UploadController) PostFetch() {
	user := services.UserTokenService.GetCurrent(this.Ctx)
	if user == nil {
		_, _ = this.Ctx.JSON(iris.Map{
			"msg":  "请先登录",
			"code": 1,
			"data": iris.Map{

			},
		})
		return
	}

	var data map[string]string
	data = make(map[string]string)

	err := this.Ctx.ReadJSON(&data)
	if err != nil {
		_, _ = this.Ctx.JSON(iris.Map{
			"msg":  err.Error(),
			"code": 0,
			"data": iris.Map{

			},
		})
		return
	}

	url := data["url"]
	output, err := oss.CopyImage(url)
	if err != nil {
		_, _ = this.Ctx.JSON(iris.Map{
			"msg":  err.Error(),
			"code": 0,
		})
	}
	_, _ = this.Ctx.JSON(iris.Map{
		"msg":  "",
		"code": 0,
		"data": iris.Map{
			"originalURL": url,
			"url":         output,
		},
	})
}
