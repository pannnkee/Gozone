package controllers

import (
	"fmt"
	"gozone/library/controller"
	"gozone/library/enum"
	"gozone/src/zone/auth"
	"gozone/src/zone/dao"
	"gozone/src/zone/model_view"
	"gozone/src/zone/models"
	"gozone/src/zone/service"
	"path"
	"time"
)

type UserController struct {
	BaseHandler
}

// 注册
func (this *UserController) Register() {
	var modelUser model_view.User
	err := controller.ParseRequestStruct(this.Controller, &modelUser)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	err, _ = new(service.RegisterService).Do(modelUser.UserName, modelUser.Email, modelUser.PassWord, modelUser.RepeatPassword, modelUser.VerifyCode)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, this.GetString("next"))
}

// 登录
func (this *UserController) Login() {
	//TODO 检查账号密码合法性
	var modelUser model_view.User
	err := controller.ParseRequestStruct(this.Controller, &modelUser)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	if modelUser.Email == "" {
		this.Response(1,"请填写登录邮箱")
	}
	if modelUser.PassWord == "" {
		this.Response(1,"请填写登录密码")
	}

	m, err := new(service.LoginService).Do(modelUser.Email, modelUser.PassWord)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	userInfo, _ := dao.UserInstance.UserInfo(modelUser.Email)
	if userInfo.Avatar == "" {
		userInfo.Avatar = "/static/img/user_avatar/default_avatar.png"
	}

	now := time.Now().Unix()
	log := models.Log{
		Ip:           this.GetIP(),
		UserID:       userInfo.Id,
		UserName:     userInfo.UserName,
		LoginTime:    now,
		LoginTimeStr: time.Unix(now, 0).Format("2006-01-02 15:04:05"),
	}
	_ = dao.LogInstance.AddLoginLog(&log)

	this.SetCK(auth.ZoneToken, string(m), 168)
	this.SetSession(SessionUserKey, userInfo)
	this.Response(0, this.GetString("next"))
}

func (this *UserController) Logout() {
	this.MustLogin()
	this.DeleteCookie(auth.ZoneToken)
	this.DelSession(SessionUserKey)
	this.Redirect("/", 302)
}

func (this *UserController) AlterPassword() {
	var modelUser model_view.User
	err := controller.ParseRequestStruct(this.Controller, &modelUser)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}
	err = new(service.AlterPasswordService).Do(this.User.Email, modelUser.PassWord, modelUser.NewPassword, modelUser.RepeatPassword)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, "")
	return
}

func (this *UserController) AlterData() {

	this.MustLogin()

	file, header, err := this.GetFile("avatar")
	if err != nil {
		this.Response(enum.DefaultError, "请选择正确的图片文件")
		return
	}
	defer file.Close()

	filename := header.Filename
	filename = fmt.Sprintf("%d.png", this.User.Id)
	err = this.SaveToFile("avatar", path.Join("static\\img\\user_avatar\\", filename))
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	exmap := map[string]interface{}{
		"avatar":    path.Join("/static/img/user_avatar/", filename),
	}
	err = dao.UserInstance.Updates(this.User.Email, exmap)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	userInfo, _ := dao.UserInstance.Get(this.User.Id)
	this.SetSession(SessionUserKey, userInfo)
	this.Response(enum.DefaultSuccess,"头像上传成功")
}

func (this *UserController) VerifyCode() {

	this.Response(enum.DefaultSuccess,"")
	//var modelUser model_view.User
	//err := controller.ParseRequestStruct(this.Controller, &modelUser)
	//if err != nil {
	//	this.Response(enum.DefaultError,err.Error())
	//	return
	//}
	//if modelUser.Email == "" {
	//	this.Response(enum.DefaultError, "邮箱为空，请检查")
	//	return
	//}
	//
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//code := fmt.Sprintf("%04v", r.Int31n(1000000))
	//_ = verifycode.Add(modelUser.Email, code)
	//
	////send email
	//err = mail.SendMail(modelUser.Email, code)
	//if err != nil {
	//	this.Response(enum.DefaultError,"发送邮件失败，请稍后重试")
	//}
	//this.Response(enum.DefaultSuccess,"发送邮件成功，请注意查收（5分钟内有效）")
}
