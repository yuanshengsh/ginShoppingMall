package main

import (
	"fmt"
	boot "ginShoppingMall/bootstrap"
	"ginShoppingMall/middleware"
	"ginShoppingMall/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	// 加载中间件
	r.Use(middleware.Cors(), middleware.JwtApp())

	// 载入路由
	r = router.AppRouter(r)

	// 启动服务
	endless.DefaultReadTimeOut = time.Second * 15
	endless.DefaultWriteTimeOut = time.Second * 15

	server := endless.NewServer(boot.Config.App.Port, r)
	//server.BeforeBegin = func(addr string) {
	//	//log.Info("Actual pid is %d, listen %s", syscall.Getpid(), addr)
	//}
	//fmt.Println("aha-log-server start " + ", listen " + conf.Conf.RUN.Addr + "......")
	over := func() { fmt.Println("exit", time.Now().Format("2006-01-02 15:04:05")) }
	server.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGINT, over)
	server.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGTERM, over)
	server.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGTSTP, over)
	server.ListenAndServe()
}
