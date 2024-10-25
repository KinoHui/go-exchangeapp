package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

// GenerateJWT 生成一个基于 JWT（JSON Web Token）的令牌，
func GenerateJWT(username string) (string, error) {
	// 创建一个新的 JWT 令牌，使用 HS256 签名方法和包含的声明。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 将用户名作为声明添加到令牌中。
		"username": username,
		// 设置令牌的过期时间为当前时间加上 72 小时。
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	// 使用给定的密钥（这里是 "secret"）签署令牌，并生成已签名的字符串。
	signedToken, err := token.SignedString([]byte("secret"))

	// 返回生成的 JWT 令牌（以 "Bearer " 为前缀）以及可能出现的错误。
	return "Bearer " + signedToken, err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
