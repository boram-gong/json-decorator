package rule

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
)

type AllRuleCfg []OneRuleCfg

type OneRuleCfg struct {
	Name  string  `json:"name"`
	Rules []*Rule `json:"rules"`
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c AllRuleCfg
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	for _, v := range c {
		for _, r := range v.Rules {
			AllRule.Store(v.Name, r)
		}
	}

	// log.Println("read config file:", cfg, "successfully")
}
