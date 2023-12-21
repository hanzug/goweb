package main

import (
	"fmt"
	"go.uber.org/zap"
	"goweb/controller"
	"goweb/dao/mysql"
	"goweb/dao/redis"
	"goweb/logger"
	"goweb/pkg/snowflake"
	"goweb/routes"
	"goweb/settings"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: config/config.yaml")
		return
	}

	fmt.Println("setting init...")
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}
	fmt.Println("setting init success")

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	zap.L().Info("logger init success")
	defer zap.L().Sync()

	zap.L().Info("mysql init...")
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	zap.L().Info("mysql init success")
	defer mysql.Close()

	zap.L().Info("redis init...")
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}
	zap.L().Info("redis init success")
	defer redis.Close()

	zap.L().Info("snowflake init...")
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		zap.L().Error("init snowflake failded", zap.Error(err))
		return
	}
	zap.L().Info("snowflake init success")

	zap.L().Info("validator init...")
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Error("init validator failded", zap.Error(err))
		return
	}
	zap.L().Info("validator init success")

	r := routes.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		zap.L().Error("run server failed", zap.Error(err))
		return
	}

}
