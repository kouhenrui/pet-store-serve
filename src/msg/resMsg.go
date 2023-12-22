package msg

const (
	//权限错误
	COOKIE_NOT_EXIST_ERROR           = "请将cookie加入请求中"
	NO_AUTH_ERROR                    = "请登录"
	ACCOUNT_PWD_ERROR                = "账号密码错误"
	SUCCESS                          = "成功"
	MODIFICATION_SUCCESSE            = "修改成功"
	ADD_SUCCESS                      = "增加成功"
	DELETE_SUCCESS                   = "删除成功"
	NAME_AND_ACCOUNT_EXIST_ERROR     = "账号已存在"
	NAME_EXIST_ERROR                 = "名称已存在"
	ACCOUNT_EXIST_ERROR              = "账号已存在"
	AUTH_LOGIN_PASSWORD_ERROR        = "密码验证错误"
	AUTH_LOGIN_COUNT_ERROR           = "账号输入错误"
	METHOD_NOT_FOUND                 = "方法未找到"
	REQUEST_NOT_ONLY_ERROR           = "请求参数详情不唯一"
	REQUEST_NOT_EXIST                = "请求参数不存在"
	AUTH_LOGIN_ERROR                 = "登陆错误"
	NO_PERMMISSION_ERROR             = "权限不足"
	REQUEST_METHOD_NOT_ALLOWED_ERROE = "请求方法不允许"
	RESOURCE_NOT_FOUND_ERROR         = "资源未找到"
	INTERNAL_ERROR                   = "服务器内部错误"
	JWT_ERROR                        = "token解析错误"
	BAD_REQUEST_ERROR                = "参数解析错误"
	PASSWORD_ENCRYPTION_ERROR        = "密码加密错误"
	PASSWORD_RESOLUTION_ERROR        = "密码解析错误"
	REDIS_INFORMATION_ERROR          = "redis数据错误"
	JSON_UNMARSHAL_ERROR             = "json反序列化失败"
	JSON_MARSHAL_ERROR               = "json序列化失败"
	ACCOUT_NOT_EXIST_ERROR           = "账号不存在"
	ADD_ERROR                        = "添加错误"
	NO_AUTHORIZATION                 = "token不能为空"
	MAKE_TOKEN_ERROR                 = "token生成错误"
	INSUFFICIENT_PERMISSION_ERROR    = "权限不足"
	METHOD_NOT_FILLED_ERROR          = "方法未填写"
	INSET_USER_ERROR                 = "写入用户错误"
	TWO_PWD_MATCH_ERROR              = "验证与原密码不一致"
	PWD_CHANGE_SUCCESS               = "密码修改成功"
	//上传接口返回值
	FILE_TYPE_ERROR   = "文件类型错误"
	FILE_SUFFIX_ERROR = "文件后缀类型不符"
	FILE_TOO_LARGE    = "上传文件太大"
	FILE_SAVE_ERROR   = "文件保存错误"
	//websocket错误

	WEBSOCKET_CONNECT_ERROR = "websocket连接错误"

	//权限错误
	PERMISSION_NOT_FOUND_ERROR = "请求权限未找到"
	PERMISSION_ADD_SUCCESS     = "权限写入成功"

	AUTHENTICATION_FAILED = "身份验证未通过"

	//图片验证码
	CAPTCHA_ERROR        = "图片验证码已过期"
	VERIFY_CODE_ERROR    = "图片验证码错误"
	CAPTCHA_CREATE_ERROR = "图片验证码生成错误"

	SQL_NOT_EXIT_ERROR = "查询记录不存在"

	//sql
	ACCOUNT_PHONE_NOT_NULL = "账号未填写"
	NOT_FOUND_ERROR        = "未找到匹配数据"

	//公共方法
	PARAM_ERROR = "参数类型错误"

	LIMITER_ERROR = "请求过于频繁，请稍后重试"
)
