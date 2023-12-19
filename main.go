package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goweb/dao/mysql"
	"goweb/dao/redis"
	"goweb/logger"
	"goweb/routes"
	"goweb/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}

	if err := logger.Init(); err != nil {
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
	redis.Close()

	r := routes.Setup()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("server and listen error", zap.String("addr", srv.Addr))
			return
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("server shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exited")
}
