package logic

import (
	"goweb/dao/mysql"
	"goweb/models"
	"goweb/pkg/jwt"
	"goweb/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {

	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
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

func Login(p *models.ParamLogin) (user *models.User, err error) {

	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	err = mysql.Login(user)

	if err != nil {
		return
	}

	// generate token

	user.Token, _ = jwt.GenToken(user.UserID, user.Username)

	return
}
