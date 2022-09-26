package service

import (
	"bullet-screen/common"
	"bullet-screen/util"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"time"
)



// 获取验证码，返回 uuid, b64s, error
func GetCaptcha() (string, string, int){

	id, b64s, code := CreateCaptcha()

	if code != util.OK {
		return "", "", code
	}


	uuid := util.GetUUID()
	code = SaveCaptcha(uuid, id)
	if code != util.OK {
		return "", "", code
	}

	return uuid, b64s, code
}

// 创建验证码
func CreateCaptcha() (string, string, int) {
	var store = base64Captcha.DefaultMemStore
	var driver base64Captcha.Driver
	var captType string


	switch captType {

	case "audio": //语音
		driver = base64Captcha.DefaultDriverAudio

	// 一些需要设置参数的验证码 https://captcha.mojotv.cn/
	//case "string": //字符验证码
		//driver = base64Captcha.NewDriverString()
	//case "math":
		//driver = base64Captcha.NewDriverMath()
	//case "chinese":
	//	driver = base64Captcha.NewDriverChinese()

	default:
		driver = base64Captcha.DefaultDriverDigit // 数字验证码

	}

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()

	// 验证码生成错误
	if err != nil {
		return "", "", util.CAPTCHA_CREATE_ERROR
	}

	return id, b64s, util.OK
}

// 将验证码 uuid-id 存储到redis中
func SaveCaptcha(key string, value string) int {
	err := common.Global_Redis.Set(key, value, time.Minute * 3).Err()
	if err != nil {
		return util.CAPTCHA_SAVE_ERROR
	}
	return util.OK
}

// 去redis中通过uuid查找对应的验证码id
func FindCaptchaByUUID(uuid string) (string, int) {
	id := common.Global_Redis.Get(uuid).Val()

	if id == "" {
		return "", util.NO_CAPTCHA
	}

	return id, util.OK
}

// 删除验证码id
func DeleteCaptchaID(uuid string) {
	err := common.Global_Redis.Del(uuid).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}


// 校验验证码, true代表校验成功, false代表校验失败
func Verify(uuid string, value string) (bool, int) {
	id, code := FindCaptchaByUUID(uuid)
	if code != util.OK {
		return false, code
	}

	if GetAnswer(id) != value {
		return false, util.CAPTCHA_WRONG
	}

	DeleteCaptchaID(uuid)
	return true, util.OK
}

// 获取验证码答案
func GetAnswer(codeID string) string {
	return base64Captcha.DefaultMemStore.Get(codeID, false)
}