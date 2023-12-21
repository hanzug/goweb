package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"goweb/models"
)

func CheckUserExits(username string) (bool, error) {
	sqlStr := "select count(user_id) from user where username = ?"

	var count int

	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser(user *models.User) (err error) {

	zap.L().Error("orginal password", zap.String("password", user.Password))
	user.Password, _ = encryptPassword(user.Password)
	zap.L().Error("now password", zap.String("password", user.Password))

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

func Login(user *models.User) (err error) {

	opassword := user.Password

	sqlStr := "select user_id, username, password from user where username = ?"

	zap.L().Info("orginal password", zap.String("password", opassword))
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		zap.L().Info("user not exist")
		return errors.New("user not exist")
	}

	zap.L().Info("now password", zap.String("password", user.Password))

	if err != nil {
		zap.L().Error("mysql err", zap.Error(err))
		return
	}

	ok := comparePasswords(user.Password, opassword)
	if !ok {
		zap.L().Info("password error")
		return errors.New("password error")
	}
	return
}
