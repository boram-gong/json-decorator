package cfg

import (
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/json-decorator/cfg/yaml"
)

var Cfg = &Config{
	PostgresConfig: &dbt.CfgDB{},
}

type Config struct {
	PostgresConfig *dbt.CfgDB
}

func InitConf(cfgPath string) {
	reader := yaml.NewYamlReader(cfgPath)
	if err := reader.ScanKey("postgres", Cfg.PostgresConfig); err != nil {
		panic(err)
	}
}
