package mysql

import (
	"golang.org/x/crypto/bcrypt"
	"goweb/models"
)

const secret = "haria"

func CheckUserExits(username string) (bool, error) {
	sqlStr := "select count(user_id) from user where username = ?"

	var count int

	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser(user *models.User) (err error) {

	user.Password, _ = encryptPassword(user.Password)

	sqlStr := "insert into user(user_id, username, password) values (?,?,?)"

	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) (string, error) {
	// 生成随机的盐值，Cost 值越大，计算哈希的时间越长
	salt, err := bcrypt.GenerateFromPassword([]byte(oPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(salt), nil
}

func comparePasswords(hashedPassword, oPassword string) bool {
	// 将字符串类型的哈希密码转换为字节切片
	hashedPasswordBytes := []byte(hashedPassword)

	// 使用 bcrypt 检查密码是否匹配
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(oPassword))
	return err == nil
}
