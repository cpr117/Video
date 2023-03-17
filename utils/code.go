// @User CPR
package utils

import (
	"VideoWeb/config"
	"math/rand"
	"time"
)

const (
	CODE_UTIL_ERR_PREFIX = "utils/code.go -> "
	CODE_REDIS_KEY       = "code"
)

func SendRegisterCode(username, email string) bool {
	code := RandomCode(config.Cfg.Code.Length)

	if SetCode(code, CODE_REDIS_KEY+username+email) {
		// 发送邮件
		go SendEmail(email, "注册验证码", "您的注册验证码为: "+code)
		return true
	}
	return true
}

func SetCode(randomCode, key string) bool {
	// 存储code到redis
	// time-to-live
	duration := Redis.TTL(key)
	if duration > time.Minute*time.Duration(config.Cfg.Code.Length-1) {
		// 还未过期
		return false
	}
	Redis.Set(key, randomCode, time.Minute*time.Duration(config.Cfg.Code.Length))
	return true
}

func RandomCode(length int) string {
	var letters = []byte("1234567890")
	result := make([]byte, length)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func VerifyCode(code, username, email string) bool {
	// 从redis中获取code
	// 比较code
	if len(code) == 0 {
		Logger.Error(CODE_UTIL_ERR_PREFIX + "code is empty")
		return false
	}
	key := CODE_REDIS_KEY + username + email
	return Redis.GetVal(key) == code
}
