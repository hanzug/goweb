package logic

import (
	"errors"
	"go.uber.org/zap"
	"goweb/dao/mysql"
	"goweb/models"
	"goweb/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {

	exist, err := mysql.CheckUserExits(p.Username)
	if err != nil {
		zap.L().Error("SignUp failed, database insert failed")
		return err
	}
	if exist {
		zap.L().Error("SignUp failed, user existed")
		return errors.New("SignUp failed, user existed")
	}

	userID := snowflake.GenID()

	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	err = mysql.InsertUser(user)
	return
}

func Login(p *models.ParamLogin) (err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
