package svc

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func MakeDebugHandler(addr string) *gin.Engine {
	engine := gin.Default()
	engine.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("http service", healthcheck.TCPDialCheck(addr, time.Second))
	engine.Handle("GET", "/health", gin.WrapF(health.LiveEndpoint))
	pprof.Register(engine)
	return engine
}
