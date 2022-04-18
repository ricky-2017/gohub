package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// 设置验证码
func (s *RedisStore) Set(key string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 本地环境方便调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

// 获取验证码值
func (s *RedisStore) Get(key string, is_clear bool) string {
	value := s.RedisClient.Get(s.KeyPrefix + key)
	if is_clear {
		s.RedisClient.Del(s.KeyPrefix + key)
	}
	return value
}

// 校验验证码
func (s *RedisStore) Verify(key string, answer string, is_clear bool) bool {
	value := s.RedisClient.Get(s.KeyPrefix + key)
	if is_clear {
		s.RedisClient.Del(s.KeyPrefix + key)
	}
	return value == answer
}
