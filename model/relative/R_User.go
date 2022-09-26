package relative

import "time"

type Regist_User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Login_User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Info_User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Gender int `json:"gender"`
	Birthday time.Time`json:"birthday"`
	Sign string `json:"sign"`
}

type UpdatePwd_User struct {
	Email string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}