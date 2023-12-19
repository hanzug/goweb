package main

import (
	"fmt"
	"go.uber.org/zap"
	"goweb/dao/mysql"
	"goweb/dao/redis"
	"goweb/logger"
	"goweb/settings"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}

	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()

	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}
	zap.L().Info("redis is running")
	redis.Close()

	//r := routes.Setup()
	//
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
	//	Handler: r,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		zap.L().Error("server and listen error", zap.String("addr", srv.Addr))
	//		return
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//zap.L().Info("Shutdown Server...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := srv.Shutdown(ctx); err != nil {
	//	zap.L().Error("server shutdown: ", zap.Error(err))
	//}
	//
	//zap.L().Info("Server exited")
}
