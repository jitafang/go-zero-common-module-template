package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 生成密码
func HashPassword(password string) string {
	// 使用默认成本（10）生成加密哈希
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// VerifyPassword 验证加密后的哈希与明文密码
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
