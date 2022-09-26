package util

const (
	OK = 200
	ERROR = 500

	// 用户问题
	WRONG_EMAIL = 2001
	WRONG_PASSWORD = 2002
	WRONG_CAPTCHA = 2003
	EXIST_EMAIL = 2004
	NOT_EXIST_EMAIL = 2005
	HAVE_NO_USERS = 2006

	// token问题
	TOKEN_IS_NULL = 3001
	TOKEN_IS_BLACKLIST = 3002
	TOKEN_IS_EXPIRES = 3003
	TOKEN_NOT_VALID = 3004
	TOKEN_MALFORMED = 3005
	TOEKN_INVALID = 3006

	// 验证码问题
	CAPTCHA_CREATE_ERROR = 4001
	CAPTCHA_SAVE_ERROR = 4002
	NO_CAPTCHA = 4003
	CAPTCHA_WRONG = 4004



	// 视频问题
	VIDEO_NOT_UPLOADED = 5001
	IN_VIDEO_PROCESSING = 5002
	UNDER_REVIEW = 5003
	PASS_THE_AUDIT = 5004
	VIDEO_CONTENT_EXIST_PROBLEM = 5005
	VIDEO_MESSAGE_EXIST_PROBLEM = 5006
	APPLY_TO_AMEND_THE_CONTENT_OF_THE_VIDEO = 5007
	APPLY_TO_AMEND_THE_MESSAGE_OF_THE_VIDEO = 5008
)

func GetMsg(code int) string {
	switch code {
	case OK: return "OK"
	case ERROR: return "ERROR"

	case WRONG_EMAIL: return "无效邮箱"
	case WRONG_PASSWORD: return "密码错误"
	case WRONG_CAPTCHA: return "验证码错误或失效"
	case EXIST_EMAIL: return "该邮箱已被使用"
	case NOT_EXIST_EMAIL: return "该邮箱未被使用"
	case HAVE_NO_USERS: return "没有用户"

	case TOKEN_IS_NULL: return "token为空，未登录或非法访问"
	case TOKEN_IS_BLACKLIST: return "账号异地登陆或令牌失效"
	case TOKEN_IS_EXPIRES: return "token授权已过期"
	case TOKEN_NOT_VALID: return "token未被激活"
	case TOKEN_MALFORMED: return "token格式错误"
	case TOEKN_INVALID: return "token不可用"

	case CAPTCHA_CREATE_ERROR: return "验证码生成错误"
	case CAPTCHA_SAVE_ERROR: return "验证码保存错误"
	case NO_CAPTCHA: return "验证码失效"
	case CAPTCHA_WRONG: return "验证码错误"



	case VIDEO_NOT_UPLOADED: return "视频未上传"
	case IN_VIDEO_PROCESSING: return "视频处理中"
	case UNDER_REVIEW: return "视频审核中"
	case PASS_THE_AUDIT: return "视频审核通过"
	case VIDEO_CONTENT_EXIST_PROBLEM: return "视频内容存在问题"
	case VIDEO_MESSAGE_EXIST_PROBLEM: return "视频信息存在问题"
	case APPLY_TO_AMEND_THE_CONTENT_OF_THE_VIDEO: return "作者申请修改视频内容"
	case APPLY_TO_AMEND_THE_MESSAGE_OF_THE_VIDEO: return "作者申请修改视频信息"

	}
	return "未知错误"
}