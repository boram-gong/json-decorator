package main

import (
	"fmt"
	dbt_pg "github.com/boram-gong/db_tool/pg"
	"github.com/boram-gong/json-decorator/cfg"
	"github.com/boram-gong/json-decorator/rule"
	"github.com/boram-gong/json-decorator/service/svc/server"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"time"
)

var (
	configFilename = kingpin.Flag("conf", "yaml config file name").Short('c').
		Default("./conf/config.yaml").String()
)

func main() {
	kingpin.Version("v0.0.1")
	kingpin.Parse()
	if _, err := os.Stat(*configFilename); err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("%s not found", *configFilename))
		} else {
			panic(err)
		}
	}
	cfg.InitConf(*configFilename)
	rule.PClient = dbt_pg.NewPgClient(cfg.Cfg.PostgresConfig)
	rule.ReAllRule()
	go func() {
		for {
			time.Sleep(10 * time.Minute)
			rule.ReAllRule()
		}
	}()
	server.Run()
}
