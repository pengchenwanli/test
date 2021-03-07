package bcrypt

import "golang.org/x/crypto/bcrypt"

//密码加密计算密码哈希值， PASSEWORD_DEFAULT表示使用bcrypt算法， cost——工作因子，表示计算次数
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//密码对对比操作
func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil //为空则加密成功
}
