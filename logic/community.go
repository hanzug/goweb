package logic

import (
	"goweb/dao/mysql"
	"goweb/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
