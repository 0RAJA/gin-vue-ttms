package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"ttms/internal/global"
	"ttms/internal/routing/router"
	"ttms/internal/setting"

	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func initSettings() {
	setting.AllInit()
}

// @title        ttms
// @version      1.0
// @description  ttms影院管理系统

// @license.name  raja,moonman
// @license.url

// @host      ttms.humraja.xyz
// @BasePath  /ttms

// @securityDefinitions.basic  BasicAuth
func main() {
	initSettings()
	if global.Settings.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.Settings.Server.Address,
		Handler:        r,
		ReadTimeout:    global.Settings.Server.ReadTimeout,
		WriteTimeout:   global.Settings.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	fmt.Println("AppName:", global.Settings.App.Name, "Version:", global.Settings.App.Version, "Address:", global.Settings.Server.Address, "RunMode:", global.Settings.Server.RunMode)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			global.Logger.Info(err.Error())
		}
	}()
	setting.Group.Auto.Init() // 需要在其他服务启动后启动自动执行

	gracefulExit(s) // 优雅退出
	global.Logger.Info("Server exited!")
}

// 优雅退出
func gracefulExit(s *http.Server) {
	// 退出通知
	quit := make(chan os.Signal, 1)
	// 等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("ShutDown Server...")
	// 给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Server.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { // 优雅退出
		global.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
	}
}
