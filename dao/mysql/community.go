package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"goweb/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"

	db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no community")
			err = nil
		}
	}
	return communityList, err
}
