package server

import (
	"fmt"
	"github.com/boram-gong/json-decorator/service/handlers"
	"github.com/boram-gong/json-decorator/service/svc"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewEndpoints() svc.Endpoints {
	endpoints := svc.Endpoints{
		JsonDecoratorEndpoint: handlers.JsonDecorator,
		ReCfgEndpoint:         handlers.ReCfg,
	}
	return endpoints
}

func interruptHandler(ch chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)
	ch <- terminateError
}

func Run() {
	//gin.DisableConsoleColor()
	//writer, _ := logs.New(
	//	cfg.Cfg.LogPath + "%Y%m%d.log",
	//)
	//
	//gin.DefaultWriter = io.MultiWriter(writer)
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	engine.Use(handlers.CustomizedMiddleware()) // 中间件

	ch := make(chan error)
	go interruptHandler(ch)

	endpoints := NewEndpoints()

	// Debug listener.
	//go func() {
	//	e := svc.MakeDebugHandler(":39999")
	//	ch <- e.Run(":39999")
	//}()

	// http
	go func() {
		svc.MakeHTTPHandler(engine, endpoints)
		ch <- engine.Run(":29989")
	}()

	fmt.Printf("closed:%s", <-ch)
}
