package service

import (
	"bullet-screen/common"
	"bullet-screen/model"
	"bullet-screen/model/public"
	"bullet-screen/model/relative"
	"bullet-screen/util"
	"fmt"
)

// 模糊查询
func SearchLike(keyword string) ([]*model.User, int) {
	keyword = util.ConsistString("%", keyword, "%")
	var users []*model.User
	err := common.Global_Mysql.Debug().Where("name like ?", keyword).Find(&users).Error
	if err != nil {
		return nil, util.HAVE_NO_USERS
	}
	return users, util.OK
}

// 修改用户密码
func UpdatePassword(u model.User, newPassword string) int {
	hash, code := GetPasswordById(u.ID)
	if code != util.OK {
		return code
	}
	if !util.BcryptCheck(u.Password, hash) {
		return util.WRONG_PASSWORD
	}

	err := common.Global_Mysql.Debug().Table("user").Where("id = ? ", u.ID).Update("password", newPassword).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}

// 通过id 获取密码
func GetPasswordById(id uint) (string, int) {
	var u model.User
	err := common.Global_Mysql.Debug().Where("id = ?", id).First(&u).Error
	if err != nil { return "", util.ERROR }

	return u.Password, util.OK
}


// 获取用户列表信息
func GetUserList(page public.Page) ([]*model.User, int) {
	var users []*model.User

	if page.PageNum <= 0 { // 获取第pageNum页的数据
		page.PageNum = 1
	}
	if page.PageSize < 0 { // 获取pageSize条数据
		page.PageSize = 10
	}

	pageCount := page.PageSize * (page.PageNum - 1)

	err := common.Global_Mysql.Debug().Limit(page.PageSize).Offset(pageCount).Find(&users).Error
	if err != nil {
		return nil, util.HAVE_NO_USERS
	}

	return users, util.OK
}

// 根据ID删除用户
func DeleteUserById(id uint) int {
	var u model.User

	err := common.Global_Mysql.Debug().Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}


// 根据ID修改部分个人信息
func ModifyInfoById(u model.User) int {
	info := make(map[string]interface{})
	info["Name"] = u.Name
	info["Gender"] = u.Gender
	info["Birthday"] = u.Birthday
	info["Sign"] = u.Sign

	err := common.Global_Mysql.Debug().Table("user").Where("id = ?", u.ID).Updates(info).Error
	if err != nil {
		fmt.Println(err)
		return util.ERROR
	}
	return util.OK
}

// 根据ID获取个人信息
func GetInfoById(id uint) (model.User, int) {
	var user model.User

	err := common.Global_Mysql.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, util.ERROR
	}

	return user, util.OK
}

// 获取用户ID
func GetIdByEmail(email string) (uint, int) {
	var user model.User

	err := common.Global_Mysql.Debug().Where("email = ?", email).First(&user).Error

	if err != nil {
		return 0, util.NOT_EXIST_EMAIL
	}

	if user.ID == 0 {
		return 0, util.NOT_EXIST_EMAIL
	}
	return user.ID, util.OK
}

// 用户登录
func Login(u model.User) int {
	b, code := CheckUserByEmail(u.Email)
	if !b {
		return code
	}

	var user model.User
	common.Global_Mysql.Debug().First(&user, "email = ?", u.Email)

	if !util.BcryptCheck(u.Password, user.Password) {
		return util.WRONG_PASSWORD
	}
	return util.OK
}

// 登录后创建token
func CreateTokenAfterLogin(u relative.Login_User) (string, error) {
	baseClaim := relative.BaseClaims{
		Email: u.Email,
	}
	jwt := util.NewJWT()

	claim := jwt.CreateClaims(baseClaim)
	token, err := jwt.CreateToken(claim)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 添加用户
func AddUser(u model.User) int {
	b, code := CheckUserByEmail(u.Email)
	if b {
		return code
	}

	u.Password = util.BcryptHash(u.Password)

	err := common.Global_Mysql.Debug().Create(&u).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}

// 判断用户是否存在（根据邮箱），如果存在该用户，返回true，否则返回false
func CheckUserByEmail(email string) (bool, int) {
	var user model.User
	err := common.Global_Mysql.Debug().Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, util.ERROR
	}

	if user.Email == "" {
		return false, util.NOT_EXIST_EMAIL
	}
	return true, util.EXIST_EMAIL
}