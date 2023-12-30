package mysql

import (
	"fmt"
	"go.uber.org/zap"
	setting "goweb/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *setting.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	zap.L().Info("unc Init(cfg *setting.MySQLConfig) (err error)", zap.Any("dsn", dsn))
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Info("sqlx.Connect(\"mysql\", dsn)", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
