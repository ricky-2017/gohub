package verifycode

type Store interface {
	// 设置验证码
	Set(id string, value string) bool

	// 获取验证码
	Get(id string, is_clear bool) string

	// 校验验证码
	Verify(id string, answer string, clear bool) bool
}
