package controller

import (
	"bullet-screen/model"
	"bullet-screen/model/public"
	"bullet-screen/model/relative"
	"bullet-screen/service"
	"bullet-screen/util"
	"github.com/gin-gonic/gin"
)

// 找回密码
func FindPassword(c *gin.Context) {
	email := c.PostForm("email")

	id, code := service.GetIdByEmail(email)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}

	util.ResponseData(code, id, c)
}

// 模糊查询
func FindLike(c *gin.Context) {
	var keyword = c.Query("keyword")
	users, code := service.SearchLike(keyword)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}

	ans := make([]*relative.Info_User, len(users))
	for i, user := range users {
		ans[i] =  &relative.Info_User{
			Name:     user.Name,
			Email:    user.Email,
			Sign:     user.Sign,
			Gender:   user.Gender,
			Birthday: user.Birthday,
		}
	}

	util.ResponseData(code, ans, c)
}

// 修改密码
func UpdatePassword(c *gin.Context) {
	var u relative.UpdatePwd_User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}

	id, code := service.GetIdByEmail(u.Email)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}

	user := model.User{
		Email: u.Email,
		Password: u.OldPassword,
	}
	user.ID = id

	code = service.UpdatePassword(user, util.BcryptHash(u.NewPassword))

	if code != util.OK { // 密码错误
		util.ResponseData(code, nil, c)
		return
	}
	util.ResponseData(code, nil ,c)
}

// 查询全部用户信息
func GetUserList(c *gin.Context)  {
	var page public.Page
	size, err := util.StringToInt(c.Query("pageSize"))
	if err != nil {
		util.ResponseData(size, nil, c)
		return
	}
	page.PageSize = size
	num, err := util.StringToInt(c.Query("pageNum"))
	if err != nil {
		util.ResponseData(num, nil, c)
		return
	}
	page.PageNum = num

	users, code := service.GetUserList(page)

	util.ResponseData(code, users, c)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var u relative.Login_User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}

	id, code := service.GetIdByEmail(u.Email)

	if code != util.OK {
		util.ResponseMap(code, nil, c)
		return
	}

	userInfo, code := service.GetInfoById(id)
	if !util.BcryptCheck(u.Password, userInfo.Password) { // 密码不匹配，不允许删除用户
		util.ResponseData(util.WRONG_PASSWORD, nil, c)
		return
	}

	code = service.DeleteUserById(id)

	util.ResponseData(code, nil, c)
}


// 修改部分个人信息(根据id)
func ModifyInfo(c *gin.Context) {
	var u relative.Info_User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}

	id, code := service.GetIdByEmail(u.Email)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}

	user := model.User{
		Name: u.Name,
		Gender: u.Gender,
		Birthday: u.Birthday,
		Sign: u.Sign,
	}
	user.ID = id

	code = service.ModifyInfoById(user)

	util.ResponseData(code, nil, c)
}

// 获取用户个人信息(根据邮箱)
func GetInfo(c *gin.Context) {
	var u relative.Login_User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}


	id, code := service.GetIdByEmail(u.Email)

	if code != util.OK {
		util.ResponseMap(code, nil, c)
		return
	}

	user, code := service.GetInfoById(id)

	responseUser := relative.Info_User{
		Email: user.Email,
		Name: user.Name,
		Gender: user.Gender,
		Birthday: user.Birthday,
		Sign: user.Sign,
	}

	util.ResponseData(code, responseUser, c)
}


// 登录
func Login(c *gin.Context) {
	var u relative.Login_User
	err := c.ShouldBindJSON(&u)
	if err != nil { return }

	user := model.User{
		Email: u.Email,
		Password: u.Password,
	}

	code := service.Login(user)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}

	token, err := service.CreateTokenAfterLogin(u)
	if err != nil {
		return
	}


	util.ResponseData(code, token, c)
}

// 注册
func Register(c *gin.Context) {
	var u relative.Regist_User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}

	user := model.User {
		Name: u.Name,
		Password: u.Password,
		Email: u.Email,
	}

	code := service.AddUser(user)

	util.ResponseMap(code, nil, c)
}