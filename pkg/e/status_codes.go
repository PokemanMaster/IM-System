package e

const (
	SUCCESS                 = 200 // SUCCESS 成功
	UPDATE_PASSWORD_SUCCESS = 201 // UPDATE_PASSWORD_SUCCESS 更新密码成功
	NOT_EXIST_IDENTIFIER    = 202 // 标识不存在
	ERROR                   = 500 // 一般性错误
	INVALID_PARAMS          = 400 // 无效参数
	ERROR_CODE              = 401 // 验证码错误

	ERROR_EXIST_NICK           = 10001 // 用户昵称已存在
	ERROR_EXIST_USER           = 10002 // 用户已存在
	ERROR_NOT_EXIST_USER       = 10003 // 用户不存在
	ERROR_NOT_COMPARE          = 10004 // 不匹配的条件
	ERROR_NOT_COMPARE_PASSWORD = 10005 // 不匹配的密码
	ERROR_FAIL_ENCRYPTION      = 10006 // 加密失败
	ERROR_NOT_EXIST_PRODUCT    = 10007 // 产品不存在
	ERROR_NOT_EXIST_ADDRESS    = 10008 // 地址不存在
	ERROR_EXIST_FAVORITE       = 10009 // 收藏已存在
	ERROR_CREATED_USER         = 10010 // 创建用户失败
	ERROR_MATCHED_USERNAME     = 10011 // 用户名必须为 6-12 位英文和数字
	ERROR_MATCHED_PASSWORD     = 10012 // 长度6-12位且包含字母或数字
	ERROR_MATCHED_TELEPHONE    = 10013 // 请输入正确的手机号

	ERROR_AUTH_CHECK_TOKEN_FAIL       = 20001 // 鉴权失败，检查Token失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT    = 20002 // Token超时
	ERROR_AUTH_TOKEN                  = 20003 // Token错误
	ERROR_AUTH                        = 20004 // 鉴权失败
	ERROR_AUTH_INSUFFICIENT_AUTHORITY = 20005 // 权限不足
	ERROR_READ_FILE                   = 20006 // 读取文件失败
	ERROR_SEND_EMAIL                  = 20007 // 发送邮件失败
	ERROR_CALL_API                    = 20008 // 调用API失败
	ERROR_UNMARSHAL_JSON              = 20009 // 反序列化JSON失败
	ERROR_DATABASE                    = 30001 // 数据库错误
	ERROR_INVALID_PASSWORD            = 40002 // 不能输入和之前一样的密码

	ERROR_OSS = 40001 // 对象存储服务错误
)

//200 OK - [GET]：服务器成功返回用户请求的数据，该操作是幂等的（Idempotent）。
//201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功。
//202 Accepted - [*]：表示一个请求已经进入后台排队（异步任务）
//204 NO CONTENT - [DELETE]：用户删除数据成功。
//400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
//401 Unauthorized - [*]：表示用户没有权限（令牌、用户名、密码错误）。
//403 Forbidden - [*] 表示用户得到授权（与401错误相对），但是访问是被禁止的。
//404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
//406 Not Acceptable - [GET]：用户请求的格式不可得（比如用户请求JSON格式，但是只有XML格式）。
//410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的。
//422 Unprocesable entity - [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误。
//500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。
