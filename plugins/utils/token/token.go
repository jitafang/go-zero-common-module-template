package token

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"plugins/constant/const_token"
	"strconv"
	"time"
)

type Token struct {
	Issuer string
	Secret string
	Exp    int
	RefExp int
}

// GetUserRedisKey  获取用户缓存Key
func GetUserRedisKey(id int64) string {
	return const_token.UserTokenKeyPrefix + strconv.FormatInt(id, 10)
}

func SaveToken(conn *redis.Redis, id int64, t Token) (string, string) {
	now := time.Now()
	token, err := generateJwtToken(t, id)
	if err != nil {
		panic(err)
	}

	key := GetUserRedisKey(id)
	duration := time.Duration(t.RefExp) * time.Hour
	refreshToken := now.Add(duration).Format(time.DateTime)
	tokenKey := key + ":" + const_token.TokenKey
	refreshTokenKey := key + ":" + const_token.RefreshTokenKey

	err = conn.Setex(tokenKey, token, t.Exp*3600)
	if err != nil {
		panic(err)
	}
	err = conn.Setex(refreshTokenKey, refreshToken, t.RefExp*3600)
	if err != nil {
		panic(err)
	}

	return token, refreshToken
}

func GetToken(conn *redis.Redis, id int64) (string, string) {
	key := GetUserRedisKey(id)
	tokenKey := key + ":" + const_token.TokenKey
	refreshTokenKey := key + ":" + const_token.RefreshTokenKey

	token, err := conn.Get(tokenKey)
	if err != nil {
		panic(err)
	}
	refreshToken, err := conn.Get(refreshTokenKey)
	if err != nil {
		panic(err)
	}
	return token, refreshToken

}

func RemoveToken(conn *redis.Redis, id int64) error {
	key := GetUserRedisKey(id)
	tokenKey := key + ":" + const_token.TokenKey
	refreshTokenKey := key + ":" + const_token.RefreshTokenKey

	_, err := conn.Del(tokenKey)
	if err != nil {
		return nil
	}
	_, err = conn.Del(refreshTokenKey)
	if err != nil {
		return err
	}
	return nil
}

func RefreshToken(conn *redis.Redis, id int64, t Token) string {
	newToken, err := generateJwtToken(t, id)
	if err != nil {
		panic(err)
	}

	key := GetUserRedisKey(id)
	tokenKey := key + ":" + const_token.TokenKey

	err = conn.Setex(tokenKey, newToken, t.Exp*3600)
	if err != nil {
		panic(err)
	}
	return newToken
}
