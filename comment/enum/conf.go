package enum


// session 配置
const (
	SessionPath         = "./temp"   // 保存路径
	SessionExpireTime   = 9600           // 有效时间,秒
	CookieExpireTime    = 1800           // 有效时间,秒
	SmsCookieExpireTime = 60             // 有效时间,秒
	LocalSessionName    = "JOKERSession" // 客户端session名称
)