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

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`

	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
