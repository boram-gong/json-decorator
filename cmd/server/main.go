package main

import (
	"github.com/boram-gong/json-decorator/rule"
	"github.com/boram-gong/json-decorator/service/svc/server"
)

func main() {
	rule.InitDB()
	rule.GetRuleCfg()
	server.Run()
}
