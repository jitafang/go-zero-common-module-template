package token

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJwtToken 生成token
func generateJwtToken(t Token, id int64) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    t.Issuer,                                                             // 签发者
		Subject:   strconv.FormatInt(id, 10),                                            // 用户id
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.Exp) * time.Hour)), // 设置过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 设置签发时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.Secret))
}

// ParseToken 解析token
func ParseToken(secret, t string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(t, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	// 验证 token 是否有效并返回声明
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
